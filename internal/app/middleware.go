package app

import (
	"context"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/xid"
	"github.com/zikaeroh/strawrank/internal/ctxlog"
	"go.uber.org/zap"
)

type pollIDKey struct{}

func (a *App) pollIDCheck(paramName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := ctxlog.FromContext(ctx)

			idStr := chi.URLParam(r, paramName)
			if idStr == "" {
				logger.Debug("empty param", zap.String("paramName", paramName))
				http.NotFound(w, r)
				return
			}

			ids, err := a.hid.DecodeInt64WithError(idStr)
			if err != nil {
				logger.Debug("error decoding pollID", zap.String("idStr", idStr), zap.Error(err))
				http.NotFound(w, r)
				return
			}

			if len(ids) == 0 {
				logger.Debug("ids had a zero length", zap.String("idStr", idStr))
				http.NotFound(w, r)
				return
			}

			ctx = context.WithValue(ctx, pollIDKey{}, ids)
			ctx, logger = ctxlog.FromContextWith(ctx, zap.Int64s("pollID", ids))

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// getPollIDs gets the poll ID slice decoded from the request. It is guaranteed
// to not be empty.
func getPollIDs(r *http.Request) []int64 {
	id := r.Context().Value(pollIDKey{})
	if id == nil {
		panic("failed to get poll ID")
	}
	return id.([]int64)
}

type userInfo struct {
	cookie xid.ID
	ip     net.IP
}

type userInfoKey struct{}

func (a *App) setUserInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := ctxlog.FromContext(ctx)

		id, err := a.getSetUserID(w, r)
		if err != nil {
			logger.Debug("failed to get user ID", zap.Error(err))
		}

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			logger.Debug("failed to split RemoteAddr", zap.Error(err), zap.String("remote_addr", r.RemoteAddr))
		}

		userIP := net.ParseIP(ip)
		if userIP == nil {
			logger.Debug("failed to parse IP", zap.Error(err), zap.String("ip", ip))
		}

		ui := userInfo{
			cookie: id,
			ip:     userIP,
		}

		ctx = context.WithValue(ctx, userInfoKey{}, ui)
		ctx, _ = ctxlog.FromContextWith(ctx, zap.String("userID", id.String()))

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func getUserInfo(r *http.Request) userInfo {
	ui := r.Context().Value(userInfoKey{})
	if ui == nil {
		panic("failed to get user info")
	}
	return ui.(userInfo)
}

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
