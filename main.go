package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
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
