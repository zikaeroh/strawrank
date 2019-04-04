package app

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/xid"
	"github.com/zikaeroh/strawrank/internal/ctxlog"
	"go.uber.org/zap"
)

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

type pollIDKey struct{}

func (a *App) pollIDCheck(paramName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := ctxlog.FromContext(ctx)

			idStr := chi.URLParam(r, paramName)
			if idStr == "" {
				logger.Debug("empty param", zap.String("paramName", paramName))
				httpError(w, http.StatusNotFound)
				return
			}

			ids, err := a.hid.DecodeWithError(idStr)
			if err != nil {
				logger.Debug("error decoding pollID", zap.String("idStr", idStr), zap.Error(err))
				httpError(w, http.StatusNotFound)
				return
			}

			// TODO: check for poll and 404 if not found (with friendly page)

			ctx = context.WithValue(ctx, pollIDKey{}, ids)
			ctx, logger = ctxlog.FromContextWith(ctx, zap.Ints("pollID", ids))

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func getPollID(r *http.Request) []int {
	id := r.Context().Value(pollIDKey{})
	if id == nil {
		panic("failed to get poll ID")
	}
	return id.([]int)
}

type userIDKey struct{}

func (a *App) getSetUserID(w http.ResponseWriter, r *http.Request) (xid.ID, error) {
	const cookieName = "user-id"

	// Using a struct with the data inside for forwards compatibility.
	cookie := struct {
		XID xid.ID `json:"xid"`
	}{}

	c, err := r.Cookie(cookieName)

	if err == nil {
		if err := a.sc.Decode(cookieName, c.Value, &cookie); err != nil {
			return xid.NilID(), err
		}

		if !cookie.XID.IsNil() {
			return cookie.XID, nil
		}
	} else if err != http.ErrNoCookie {
		return xid.NilID(), err
	}

	cookie.XID = xid.New()

	encoded, err := a.sc.Encode(cookieName, cookie)
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

func (a *App) userIDCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := ctxlog.FromContext(ctx)

		user, err := a.getSetUserID(w, r)
		if err != nil {
			logger.Debug("failed to get user ID", zap.Error(err))
			httpError(w, http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, userIDKey{}, user)
		ctx, _ = ctxlog.FromContextWith(ctx, zap.String("userID", user.String()))

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func getUserID(r *http.Request) xid.ID {
	id := r.Context().Value(userIDKey{})
	if id == nil {
		panic("failed to get user ID")
	}
	return id.(xid.ID)
}

type requestIDKey struct{}

const requestIDHeader = "X-Request-ID"

// TODO: move these functions into a package

func injectRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		requestID := r.Header.Get(requestIDHeader)
		id, err := xid.FromString(requestID)
		if err != nil {
			id = xid.New()
			requestID = id.String()
		}

		w.Header().Set(requestIDHeader, requestID)
		ctx = context.WithValue(ctx, requestIDKey{}, id)

		ctx, _ = ctxlog.FromContextWith(ctx,
			zap.String("requestID", requestID),
		)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func getRequestID(r *http.Request) xid.ID {
	requestID, _ := r.Context().Value(requestIDKey{}).(xid.ID)
	return requestID
}
