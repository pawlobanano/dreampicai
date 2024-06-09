package handler

import (
	"context"
	"dreampicai/pkg/sb"
	"dreampicai/types"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
)

func WithAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}
		user := getAuthenticatedUser(r)
		if !user.IsLoggedIn {
			path := r.URL.Path
			http.Redirect(w, r, "/login?to="+path, http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func WithLogger(s types.Server) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), types.ContextKey("logger"), s.Logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func WithUser(s types.Server) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.URL.Path, "/public") {
				next.ServeHTTP(w, r)
				return
			}
			store := sessions.NewCookieStore([]byte(s.Config.SessionSecret))
			session, err := store.Get(r, s.Config.SessionUserKey)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			accessToken := session.Values[s.Config.SessionAccessTokenKey]
			if accessToken == nil {
				next.ServeHTTP(w, r)
				return
			}
			resp, err := sb.Client.Auth.User(r.Context(), accessToken.(string))
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user := types.AuthenticatedUser{
				Email:      resp.Email,
				IsLoggedIn: true,
			}
			ctx := context.WithValue(r.Context(), types.UserContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
