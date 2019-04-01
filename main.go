package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/csrf"
	"github.com/rs/zerolog"
	"github.com/zikaeroh/strawrank/internal/templates"
)

func main() {
	logger := zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = time.RFC3339
	})).With().Timestamp().Caller().Logger()

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

	r.Use(csrf.Protect([]byte("a-32-byte-long-key-goes-here"),
		csrf.Secure(false), // TODO: debug flag
	))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Route("/{id}", func(r chi.Router) {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost {
				r.ParseForm()
				spew.Dump(r.Form)
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
