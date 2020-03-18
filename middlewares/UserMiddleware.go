package middlewares

import (
	"context"
	"net/http"

	"github.com/oxodao/vibes/dal"
	"github.com/oxodao/vibes/services"
)

//CheckUserMiddleware is a middleware that add the user object to the context according to its cookie
func CheckUserMiddleware(prv *services.Provider, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("PHPSESSID")

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		u, err := dal.FindUserByToken(prv, cookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserContext, u)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}
