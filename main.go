package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/csrf"
	"github.com/gorilla/securecookie"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/zikaeroh/strawrank/internal/templates"
)

func main() {
	secureKey := []byte("a-32-byte-long-key-goes-here")

	logger := zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = time.RFC3339
	})).With().Timestamp().Caller().Logger()

	sc := securecookie.New(secureKey, nil)

	r := chi.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = logger.WithContext(ctx)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	})

	r.Use(requestLogger)
	r.Use(recoverer)

	r.Use(csrf.Protect(
		secureKey,
		csrf.Secure(false), // TODO: debug flag
	))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})

	r.Route("/{id}", func(r chi.Router) {
		fn := func(w http.ResponseWriter, r *http.Request) {
			logger := zerolog.Ctx(r.Context())

			user, err := getSetUserID(sc, w, r)
			if err != nil {
				logger.Debug().Err(err).Msg("failed to get user ID")
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			// Hack to deal with zerolog's bad API
			tmpLogger := logger.With().Str("user", user.String()).Logger()
			logger = &tmpLogger

			if r.Method == http.MethodPost {
				votesStr := r.FormValue("votes")
				var votes []int

				if err := json.Unmarshal([]byte(votesStr), &votes); err != nil {
					// TODO: Do someting in the UI instead.
					http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
					return
				}

				logger.Debug().Ints("votes", votes).Msg("posted vote")

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

		r.Get("/", fn)
		r.Post("/", fn)

		r.Get("/r", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		logger.Fatal().Err(err).Msg("exiting")
	}
}

func requestLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()

		defer func() {
			duration := time.Since(start)
			logger := zerolog.Ctx(r.Context())

			logger.Info().
				Str("method", r.Method).
				Str("url", r.RequestURI).
				Str("proto", r.Proto).
				Int("status", ww.Status()).
				Int("size", ww.BytesWritten()).
				Dur("duration", duration).
				Msg("http request")
		}()

		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(fn)
}

func recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				logger := zerolog.Ctx(r.Context())

				logger.Error().
					Stack().
					Err(errors.New("panic")).
					Interface("panic_value", rvr).
					Msg("panic")

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
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
