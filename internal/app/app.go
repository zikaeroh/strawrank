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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/zikaeroh/strawrank/internal/app/mid"
	"github.com/zikaeroh/strawrank/internal/ctxlog"
	"github.com/zikaeroh/strawrank/internal/db/models"
	"github.com/zikaeroh/strawrank/internal/static"
	"github.com/zikaeroh/strawrank/internal/templates"
	"go.uber.org/zap"
)

type Config struct {
	Logger *zap.Logger

	DB *sql.DB

	CookieKey []byte

	HIDMinLength int
	HIDSalt      string

	RealIP bool

	Debug bool
}

type App struct {
	c   *Config
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
		c:  c,
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

	r.NotFound(a.handleNotFound)

	if c.Logger != nil {
		r.Use(ctxlog.Logger(c.Logger))
	}

	r.Use(mid.RequestID)

	if c.RealIP {
		r.Use(middleware.RealIP)
	}

	r.Use(mid.RequestLogger)
	r.Use(mid.Recoverer)

	// Secure is false as this will likely be run behind a proxy.
	r.Use(csrf.Protect(c.CookieKey, csrf.Secure(false)))

	r.Get("/", a.handleIndex)
	r.Post("/", a.handleIndexPost)

	r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(static.FS(false))))
	r.Handle("/favicon.ico", http.RedirectHandler("/static/favicon.ico", http.StatusMovedPermanently))

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

func (a *App) handleNotFound(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.NotFoundPage{})
}

func (a *App) handleAbout(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.AboutPage{})
}

func (a *App) handleIndex(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.IndexPage{
		CSRF: string(csrf.TemplateField(r)),
	})
}

func (a *App) handleIndexPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		a.internalServerError(w, err)
		return
	}

	ctx := r.Context()
	logger := ctxlog.FromContext(ctx)

	question := r.FormValue("question")
	choices := r.Form["choice"]
	checkMode := r.FormValue("checkMode")

	question = strings.TrimSpace(question)

	if len(question) == 0 {
		a.badRequest(w, "empty question")
		return
	}

	if len(question) > 100 {
		a.badRequest(w, "question too long")
		return
	}

	if len(choices) <= 1 {
		a.badRequest(w, "not enough choices")
		return
	}

	for i, choice := range choices {
		choice = strings.TrimSpace(choice)

		if len(choice) == 0 {
			a.badRequest(w, "empty choice")
			return
		}

		if len(choice) > 50 {
			a.badRequest(w, "choice too long")
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
		a.badRequest(w, "bad check mode "+checkMode)
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
		a.internalServerError(w, err)
		return
	}

	p, err := a.hid.EncodeInt64([]int64{poll.ID})
	if err != nil {
		logger.Error("error encoding poll ID", zap.Int64("id", poll.ID))
		a.internalServerError(w, err)
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
		a.internalServerError(w, err)
		return
	}

	templates.WritePageTemplate(w, &templates.VotePage{
		CSRF:     string(csrf.TemplateField(r)),
		Path:     r.URL.Path,
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
		a.badRequest(w, "failed to unmarshal votes: "+err.Error())
		// logger.Debug("failed to unmarshal votes", zap.Error(err))
		return
	}

	if len(votes) == 0 {
		a.badRequest(w, "empty votes")
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
			a.internalServerError(w, err)
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
				a.internalServerError(w, err)
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
				a.badRequest(w, fmt.Sprintf("vote is out of range: %d in [0, %d)", vote, choicesLen))
				// logger.Debug("vote is out of range", zap.Int64("vote", vote), zap.Int64("len", choicesLen))
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
			a.internalServerError(w, err)
			return err
		}

		http.Redirect(w, r, r.RequestURI+"/r", http.StatusSeeOther)
		return nil
	})

	if txErr != nil {
		logger.Error("transaction error", zap.Error(txErr))
		a.internalServerError(w, txErr)
		return
	}
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
