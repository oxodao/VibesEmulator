package routes

import (
	"net/http"
)

// IndexRoute represents the index route
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The Test 1.2.1 private server"))
}