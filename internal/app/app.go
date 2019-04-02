package app

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/csrf"
	"github.com/gorilla/securecookie"
	"github.com/rs/xid"
	"github.com/zikaeroh/strawrank/internal/ctxlog"
	"github.com/zikaeroh/strawrank/internal/templates"
	"go.uber.org/zap"
)

type Config struct {
	Logger    *zap.Logger
	CookieKey []byte
}

type App struct {
	r  chi.Router
	sc *securecookie.SecureCookie
}

func New(c *Config) *App {
	a := &App{
		sc: securecookie.New(c.CookieKey, nil),
	}

	r := chi.NewRouter()
	a.r = r

	if c.Logger != nil {
		r.Use(ctxlog.Logger(c.Logger))
	}

	r.Use(requestLogger)
	r.Use(recoverer)
	r.Use(csrf.Protect(c.CookieKey, csrf.Secure(false)))

	r.Get("/", a.handleIndex)
	r.Get("/favicon.ico", http.NotFound)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", a.handleVote)
		r.Post("/", a.handleVote)
		r.Get("/r", a.handleResults)
	})

	return a
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.r.ServeHTTP(w, r)
}

func (a *App) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func (a *App) handleVote(w http.ResponseWriter, r *http.Request) {
	logger := ctxlog.FromRequest(r)

	user, err := getSetUserID(a.sc, w, r)
	if err != nil {
		logger.Debug("failed to get user ID", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	// Hack to deal with zerolog's bad API
	logger = logger.With(zap.String("user", user.String()))

	if r.Method == http.MethodPost {
		votesStr := r.FormValue("votes")
		var votes []int

		if err := json.Unmarshal([]byte(votesStr), &votes); err != nil {
			// TODO: Do someting in the UI instead.
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		logger.Debug("posted vote", zap.Ints("votes", votes))

		// Post/Redirect/Get
		http.Redirect(w, r, r.URL.String()+"/r", http.StatusSeeOther)
		return
	}

	templates.WritePageTemplate(w, &templates.VotePage{
		CSRF: string(csrf.TemplateField(r)),
		Name: "What should we do today?",
		Choices: []string{
			"A",
			"B",
			"C",
		},
	})
}

func (a *App) handleResults(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func getSetUserID(s securecookie.Codec, w http.ResponseWriter, r *http.Request) (xid.ID, error) {
	const cookieName = "user-id"

	// Using a struct with the data inside for forwards compatibility.
	cookie := struct {
		XID xid.ID `json:"xid"`
	}{}

	c, err := r.Cookie(cookieName)

	if err == nil {
		if err := s.Decode(cookieName, c.Value, &cookie); err != nil {
			return xid.NilID(), err
		}

		if !cookie.XID.IsNil() {
			return cookie.XID, nil
		}
	} else if err != http.ErrNoCookie {
		return xid.NilID(), err
	}

	cookie.XID = xid.New()

	encoded, err := s.Encode(cookieName, cookie)
	if err != nil {
		return xid.NilID(), err
	}

	c = &http.Cookie{
		Name:     cookieName,
		Value:    encoded,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, c)

	return cookie.XID, nil
}
