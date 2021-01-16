package main

import (
	"flag"
	"fmt"
	"github.com/oxodao/vibes/config"
	"github.com/oxodao/vibes/patcher"
	"log"
	"net/http"
	"os"
	"strings"
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

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	prv := services.NewProvider(cfg)

	r := mux.NewRouter()

	routes.Auth(prv, r.PathPrefix("/auth/").Subrouter())
	routes.Core(prv, r.PathPrefix("/core/").Subrouter())
	routes.Game(prv, r.PathPrefix("/game/").Subrouter())
	routes.Messenger(prv, r.PathPrefix("/messenger/").Subrouter())
	routes.Settings(prv, r.PathPrefix("/settings/").Subrouter())


	debug := true
	r.PathPrefix("/pictures/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if debug {
			fmt.Println("Access on " + r.RequestURI)
			fileName := r.URL.RequestURI()[strings.LastIndex(r.URL.RequestURI(), "/")+1:]

			if _, err := os.Stat("./pictures/" + fileName); os.IsNotExist(err) {
				fmt.Println(" CAN'T FIND PICTURE SIZE ", fileName)
			}
		}
		http.StripPrefix("/pictures/", http.FileServer(http.Dir("./pictures"))).ServeHTTP(w, r)
	})

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("UNHANDLED ACCESS: " + r.RequestURI)
	})

	srv := &http.Server{
		Handler: r,
		Addr:    cfg.ListeningAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Listening on ", srv.Addr)
	log.Fatal(srv.ListenAndServe())

}
