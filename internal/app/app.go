package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/csrf"
	"github.com/gorilla/securecookie"
	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
	"github.com/speps/go-hashids"
	"github.com/zikaeroh/strawrank/internal/app/mid"
	"github.com/zikaeroh/strawrank/internal/ctxlog"
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
	db  *sqlx.DB
	sc  *securecookie.SecureCookie
	hid *hashids.HashID
}

func New(c *Config) (*App, error) {
	if c.DB == nil {
		return nil, errors.New("db was nil") // TODO: make into a value
	}

	var err error

	a := &App{
		db: sqlx.NewDb(c.DB, "postgres"),
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
			r.Use(a.userIDCheck)
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

	logger := ctxlog.FromRequest(r)

	question := r.FormValue("question")
	choices := r.Form["choice"]

	logger.Debug("posted new poll", zap.String("question", question), zap.Strings("choices", choices))

	// TODO: store submission, redirect to results page

	p, err := a.hid.Encode([]int{1})
	if err != nil {
		httpError(w, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/p/"+p, http.StatusSeeOther)
}

func (a *App) handleVote(w http.ResponseWriter, r *http.Request) {
	// pollIDs := getPollID(r)
	// userID := getUserID(r)

	templates.WritePageTemplate(w, &templates.VotePage{
		CSRF: string(csrf.TemplateField(r)),
		Name: "What should we do today?",
		Choices: []string{
			"This is a super long choice that likely wraps and that's not so good. Let's make it even longer, shall we?",
			"B",
			"C",
		},
	})
}

func (a *App) handleVotePost(w http.ResponseWriter, r *http.Request) {
	logger := ctxlog.FromRequest(r)

	// pollIDs := getPollID(r)
	// userID := getUserID(r)

	votesStr := r.FormValue("votes")
	var votes []int

	if err := json.Unmarshal([]byte(votesStr), &votes); err != nil {
		// TODO: Do someting in the UI instead.
		httpError(w, http.StatusBadRequest)
		return
	}

	if len(votes) == 0 {
		httpError(w, http.StatusBadRequest)
		return
	}

	logger.Debug("posted vote", zap.Ints("votes", votes))

	// TODO: tally votes

	// Post/Redirect/Get
	http.Redirect(w, r, r.RequestURI+"/r", http.StatusSeeOther)

}

func (a *App) handleResults(w http.ResponseWriter, r *http.Request) {
	// pollIDs := getPollID(r)

	templates.WritePageTemplate(w, &templates.ResultsPage{
		Name: "This is a test",
	})
}

func (a *App) handleAbout(w http.ResponseWriter, r *http.Request) {
	templates.WritePageTemplate(w, &templates.AboutPage{})
}

func (a *App) handleDebugDatabase(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "polls:")

	type poll struct {
		ID       int
		Question string
		Choices  []string
	}

	rows, err := a.db.Queryx(`SELECT * FROM polls`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var p poll

		if err := rows.StructScan(&p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%#v", p)
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "ballots:")

	type ballot struct {
		ID      int
		PollID  int    `db:"poll_id"`
		UserXID xid.ID `db:"user_xid"`
		UserIP  string `db:"user_ip"`
		Votes   []int
	}

	rows, err = a.db.Queryx(`SELECT * FROM ballots`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var b ballot

		if err := rows.StructScan(&b); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "%#v", b)
	}
}
