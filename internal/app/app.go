package app

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/csrf"
	"github.com/gorilla/securecookie"
	"github.com/speps/go-hashids"
	"github.com/zikaeroh/strawrank/internal/app/mid"
	"github.com/zikaeroh/strawrank/internal/ctxlog"
	"github.com/zikaeroh/strawrank/internal/templates"
	"go.uber.org/zap"
)

type Config struct {
	Logger *zap.Logger

	CookieKey []byte

	HIDMinLength int
	HIDSalt      string
}

type App struct {
	r   chi.Router
	sc  *securecookie.SecureCookie
	hid *hashids.HashID
}

func New(c *Config) (*App, error) {
	var err error

	a := &App{
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

	r.Route("/{pollID}", func(r chi.Router) {
		r.Use(a.pollIDCheck("pollID"))

		r.Group(func(r chi.Router) {
			r.Use(a.userIDCheck)
			r.Get("/", a.handleVote)
			r.Post("/", a.handleVote)
		})

		r.With(middleware.NoCache).Get("/r", a.handleResults)
	})

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

	question := r.FormValue("question")
	choices := r.Form["choice"]

	_ = question
	_ = choices

	// TODO: store submission, redirect to results page

	_, _ = w.Write([]byte("ok"))
}

func (a *App) handleVote(w http.ResponseWriter, r *http.Request) {
	logger := ctxlog.FromRequest(r)

	// pollIDs := getPollID(r)
	// userID := getUserID(r)

	if r.Method == http.MethodPost {
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
		http.Redirect(w, r, r.URL.String()+"/r", http.StatusSeeOther)
		return
	}

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

func (a *App) handleResults(w http.ResponseWriter, r *http.Request) {
	// pollIDs := getPollID(r)

	templates.WritePageTemplate(w, &templates.ResultsPage{
		Name: "This is a test",
	})
}
