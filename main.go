package main

import (
	"flag"
	"fmt"
	"github.com/oxodao/vibes/patcher"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oxodao/vibes/routes"
	"github.com/oxodao/vibes/services"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	patchFlag := flag.String("path", "", "APK to patch")
	newURLFlag := flag.String("url", "", "New URL for the API")

	flag.Parse()

	if len(*patchFlag) > 0 && len(*newURLFlag) > 0 {
		patcher.Patcher(*patchFlag, *newURLFlag)
		return
	} else if len(*patchFlag) > 0 || len(*newURLFlag) > 0 {
		fmt.Println("Vibes APK Patcher")
		fmt.Println("In order to patch the APK you must BOTH give the -path <APK Path> and -url <new api url> arguments")

		return
	}

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
