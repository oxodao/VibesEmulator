package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oxodao/vibes/routes"
	"github.com/oxodao/vibes/services"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	fmt.Println("Vibes API - Indev")

	prv := services.NewProvider()

	r := mux.NewRouter()

	routes.Auth(prv, r.PathPrefix("/auth/").Subrouter())
	routes.Core(prv, r.PathPrefix("/core/").Subrouter())
	routes.Messenger(prv, r.PathPrefix("/messenger/").Subrouter())
	routes.Settings(prv, r.PathPrefix("/settings/").Subrouter())


	r.PathPrefix("/pictures/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Access on " + r.RequestURI)
		http.StripPrefix("/pictures/", http.FileServer(http.Dir("./pictures"))).ServeHTTP(w, r)
	})

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("UNHANDLED ACCESS: " + r.RequestURI)
	})

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:4568",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
