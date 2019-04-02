package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/csrf"
	"github.com/gorilla/securecookie"
	"github.com/rs/xid"
	"github.com/zikaeroh/strawrank/internal/ctxlog"
	"github.com/zikaeroh/strawrank/internal/templates"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TODO: flag
const debug = true

func main() {
	// TODO: flag
	secureKey := []byte("a-32-byte-long-key-goes-here")

	var logConfig zap.Config

	if debug {
		logConfig = zap.NewDevelopmentConfig()
		logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		logConfig = zap.NewProductionConfig()
	}

	logger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}

	sc := securecookie.New(secureKey, nil)

	r := chi.NewRouter()

	r.Use(ctxlog.Logger(logger))
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
			logger := ctxlog.FromRequest(r)

			user, err := getSetUserID(sc, w, r)
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

		r.Get("/", fn)
		r.Post("/", fn)

		r.Get("/r", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ok"))
		})
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		logger.Fatal("exiting",
			zap.Error(err),
		)
	}
}

func requestLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()

		defer func() {
			duration := time.Since(start)
			logger := ctxlog.FromRequest(r)

			logger.Info("http request",
				zap.String("method", r.Method),
				zap.String("url", r.RequestURI),
				zap.String("proto", r.Proto),
				zap.Int("status", ww.Status()),
				zap.Int("size", ww.BytesWritten()),
				zap.Duration("duration", duration),
			)
		}()

		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(fn)
}

func recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				logger := ctxlog.FromRequest(r)

				// Ensure logger is logging stack traces, at least here.
				logger = logger.WithOptions(zap.AddStacktrace(zap.ErrorLevel))

				logger.Error("PANIC",
					zap.Any("panic_value", rvr),
				)

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
