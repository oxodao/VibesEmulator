package middlewares

import (
	"net/http"
)

// ContentTypeMiddleware tmp
func ContentTypeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("A request to this url has been made: " + r.RequestURI)

		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
}
