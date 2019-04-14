package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/csrf"
	"github.com/gorilla/securecookie"
	"github.com/speps/go-hashids"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/zikaeroh/strawrank/internal/app/mid"
	"github.com/zikaeroh/strawrank/internal/ctxlog"
	"github.com/zikaeroh/strawrank/internal/db/models"
	"github.com/zikaeroh/strawrank/internal/templates"
	"go.uber.org/zap"
)

type Config struct {
	Logger *zap.Logger

	DB *sql.DB

	CookieKey []byte

	HIDMinLength int
	HIDSalt      string

	Debug bool
}

type App struct {
	r   chi.Router
	db  *sql.DB
	sc  *securecookie.SecureCookie
	hid *hashids.HashID
}

func New(c *Config) (*App, error) {
	if c.DB == nil {
		return nil, errors.New("db was nil") // TODO: make into a value
	}

	var err error

	a := &App{
		db: c.DB,
		sc: securecookie.New(c.CookieKey, nil),
	}

	a.hid, err = hashids.NewWithData(&hashids.HashIDData{
		Alphabet:  hashids.DefaultAlphabet,
		MinLength: c.HIDMinLength,
		Salt:      c.HIDSalt,
	})
	if err != nil {
		return nil, err
	}

	r := chi.NewRouter()
	a.r = r

	if c.Logger != nil {
		r.Use(ctxlog.Logger(c.Logger))
	}

	r.Use(mid.RequestID)
	r.Use(mid.RequestLogger)
	r.Use(mid.Recoverer)

	// Secure is false as this will likely be run behind a proxy.
	r.Use(csrf.Protect(c.CookieKey, csrf.Secure(false)))

	r.Get("/", a.handleIndex)
	r.Post("/", a.handleIndexPost)

	r.Get("/favicon.ico", http.NotFound) // TODO

	r.Get("/about", a.handleAbout)

	r.Route("/p/{pollID}", func(r chi.Router) {
		r.Use(a.pollIDCheck("pollID"))

		r.Group(func(r chi.Router) {
			r.Use(a.setUserInfo)
			r.Get("/", a.handleVote)
			r.Post("/", a.handleVotePost)
		})

		r.With(middleware.NoCache).Get("/r", a.handleResults)
	})

	if c.Debug {
		r.Route("/debug", func(r chi.Router) {
			r.Get("/database", a.handleDebugDatabase)
		})
	}

	return a, nil
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.r.ServeHTTP(w, r)
}

func (a *App) handleIndex(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.IndexPage{
		CSRF: string(csrf.TemplateField(r)),
	})
}

func (a *App) handleIndexPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		httpError(w, http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	logger := ctxlog.FromContext(ctx)

	question := r.FormValue("question")
	choices := r.Form["choice"]
	checkMode := r.FormValue("checkMode")

	question = strings.TrimSpace(question)

	if len(question) == 0 {
		logger.Debug("empty question")
		httpError(w, http.StatusBadRequest)
		return
	}

	if len(question) > 100 {
		logger.Debug("question too long")
		httpError(w, http.StatusBadRequest)
		return
	}

	if len(choices) <= 1 {
		logger.Debug("not enough choices")
		httpError(w, http.StatusBadRequest)
		return
	}

	for i, choice := range choices {
		choice = strings.TrimSpace(choice)

		if len(choice) == 0 {
			logger.Debug("empty choice")
			httpError(w, http.StatusBadRequest)
			return
		}

		if len(choice) > 50 {
			logger.Debug("choice too long")
			httpError(w, http.StatusBadRequest)
			return
		}

		choices[i] = choice
	}

	checkMode = strings.TrimSpace(checkMode)

	switch checkMode {
	case models.BallotCheckModeNone:
	case models.BallotCheckModeCookie:
	case models.BallotCheckModeIP:
	case models.BallotCheckModeIPAndCookie:
		// Valid
	default:
		logger.Debug("bad check mode", zap.String("check_mode", checkMode))
		httpError(w, http.StatusBadRequest)
		return
	}

	logger.Debug("posted new poll", zap.String("question", question), zap.Strings("choices", choices))

	poll := models.Poll{
		Question:  question,
		Choices:   choices,
		CheckMode: checkMode,
	}

	if err := poll.Insert(ctx, a.db, boil.Infer()); err != nil {
		logger.Error("error inserting poll", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	p, err := a.hid.EncodeInt64([]int64{poll.ID})
	if err != nil {
		logger.Error("error encoding poll ID", zap.Int64("id", poll.ID))
		httpError(w, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/p/"+p, http.StatusSeeOther)
}

func (a *App) handleVote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := ctxlog.FromContext(ctx)

	pollID := getPollIDs(r)[0]

	poll, err := models.FindPoll(ctx, a.db, pollID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}

		logger.Error("error finding poll", zap.Error(err))
		httpError(w, http.StatusInternalServerError)
		return
	}

	templates.WritePageTemplate(w, &templates.VotePage{
		CSRF:     string(csrf.TemplateField(r)),
		Question: poll.Question,
		Choices:  poll.Choices,
	})
}

func (a *App) handleVotePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := ctxlog.FromContext(ctx)

	pollID := getPollIDs(r)[0]
	ui := getUserInfo(r)

	var votes []int64

	if err := json.Unmarshal([]byte(r.FormValue("votes")), &votes); err != nil {
		logger.Debug("failed to unmarshal votes", zap.Error(err))
		httpError(w, http.StatusBadRequest)
		return
	}

	if len(votes) == 0 {
		logger.Debug("empty votes")
		httpError(w, http.StatusBadRequest)
		return
	}

	logger.Debug("posted vote", zap.Int64s("votes", votes))

	txErr := a.transact(func(tx *sql.Tx) error {
		poll, err := models.FindPoll(ctx, a.db, pollID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.NotFound(w, r)
				return nil
			}

			logger.Error("error finding poll", zap.Error(err))
			httpError(w, http.StatusInternalServerError)
			return err
		}

		var cookie null.String
		if !ui.cookie.IsNil() {
			cookie = null.StringFrom(ui.cookie.String())
		}

		var userIP null.String
		if len(ui.ip) != 0 {
			userIP = null.StringFrom(ui.ip.String())
		}

		var qms []qm.QueryMod

		switch poll.CheckMode {
		case models.BallotCheckModeNone:
			// Do nothing.
		case models.BallotCheckModeCookie:
			qms = []qm.QueryMod{
				models.BallotWhere.PollID.EQ(pollID),
				models.BallotWhere.Cookie.EQ(cookie),
			}
		case models.BallotCheckModeIP:
			qms = []qm.QueryMod{
				models.BallotWhere.PollID.EQ(pollID),
				models.BallotWhere.IPAddr.EQ(userIP),
			}
		case models.BallotCheckModeIPAndCookie:
			qms = []qm.QueryMod{
				models.BallotWhere.PollID.EQ(pollID),
				qm.Expr(
					models.BallotWhere.Cookie.EQ(cookie),
					qm.Or2(models.BallotWhere.IPAddr.EQ(userIP)),
				),
			}
		default:
			panic("unreachable")
		}

		if len(qms) != 0 {
			exists, err := models.Ballots(qms...).Exists(ctx, tx)
			if err != nil {
				logger.Error("error checking for existing ballot", zap.Error(err))
				httpError(w, http.StatusInternalServerError)
				return err
			}

			if exists {
				// TODO: indicate that this is a duplicate?
				http.Redirect(w, r, r.RequestURI+"/r", http.StatusSeeOther)
				return nil
			}
		}

		choicesLen := int64(len(poll.Choices))

		for _, vote := range votes {
			if vote < 0 || vote >= choicesLen {
				logger.Debug("vote is out of range", zap.Int64("vote", vote), zap.Int64("len", choicesLen))
				httpError(w, http.StatusBadRequest)
				return nil
			}
		}

		ballot := models.Ballot{
			PollID: pollID,
			Cookie: cookie,
			IPAddr: userIP,
			Votes:  votes,
		}

		if err := ballot.Insert(ctx, a.db, boil.Infer()); err != nil {
			logger.Error("error inserting ballot", zap.Error(err))
			httpError(w, http.StatusInternalServerError)
			return err
		}

		http.Redirect(w, r, r.RequestURI+"/r", http.StatusSeeOther)
		return nil
	})

	if txErr != nil {
		logger.Error("transaction error", zap.Error(txErr))
		httpError(w, http.StatusInternalServerError)
		return
	}
}

func (a *App) handleAbout(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.AboutPage{})
}

func (a *App) handleDebugDatabase(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Fprintln(w, "polls:")

	polls, err := models.Polls().All(ctx, a.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)

	for _, poll := range polls {
		if err := enc.Encode(poll); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w)
	}

	fmt.Fprintln(w, "ballots:")

	ballots, err := models.Ballots().All(ctx, a.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, ballot := range ballots {
		if err := enc.Encode(ballot); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w)
	}

	fmt.Fprintln(w, "ballots from [::1]:")

	ip := net.ParseIP("::1").To16()

	ballots, err = models.Ballots(models.BallotWhere.IPAddr.EQ(null.StringFrom(ip.String()))).All(ctx, a.db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, ballot := range ballots {
		if err := enc.Encode(ballot); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w)
	}
}
