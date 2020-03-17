package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oxodao/vibes/middlewares"
	"github.com/oxodao/vibes/models"
	"github.com/oxodao/vibes/routes"
	"github.com/oxodao/vibes/services"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	fmt.Println("Vibes API - Indev")

	prv := services.NewProvider()
	prv.DB.AutoMigrate(&models.Picture{})
	prv.DB.AutoMigrate(&models.User{})
	prv.DB.AutoMigrate(&models.Contact{})

	r := mux.NewRouter()

	auth := r.PathPrefix("/auth/").Subrouter()
	auth.HandleFunc("/register", routes.RegisterRoute(prv))

	core := r.PathPrefix("/core/").Subrouter()
	core.HandleFunc("/uploadPicture", routes.UploadPictureRoute(prv))
	core.HandleFunc("/getContacts", middlewares.CheckUserMiddleware(prv, routes.GetContactsRoute(prv)))
	//core.HandleFunc("/getPotentialContacts", middlewares.CheckUserMiddleware(prv, routes.GetPotentialContactsRoute(prv)))

	settings := r.PathPrefix("/settings/").Subrouter()
	settings.HandleFunc("/getAll", middlewares.CheckUserMiddleware(prv, routes.GetAllSettingsRoute(prv)))
	settings.HandleFunc("/setAge", middlewares.CheckUserMiddleware(prv, routes.SetAgeRoute(prv)))
	settings.HandleFunc("/setAgeRange", middlewares.CheckUserMiddleware(prv, routes.SetAgeRangeRoute(prv)))
	settings.HandleFunc("/setFirstName", middlewares.CheckUserMiddleware(prv, routes.SetFirstNameRoute(prv)))
	settings.HandleFunc("/setGameLanguage", middlewares.CheckUserMiddleware(prv, routes.SetGameLanguageRoute(prv)))
	settings.HandleFunc("/setGender", middlewares.CheckUserMiddleware(prv, routes.SetGenderRoute(prv)))
	settings.HandleFunc("/setGenderWanted", middlewares.CheckUserMiddleware(prv, routes.SetGenderWantedRoute(prv)))
	settings.HandleFunc("/setPicture", middlewares.CheckUserMiddleware(prv, routes.SetPictureRoute(prv)))
	settings.HandleFunc("/setPushToken", middlewares.CheckUserMiddleware(prv, routes.SetPushTokenRoute(prv)))
	settings.HandleFunc("/setXRatedEnabled", middlewares.CheckUserMiddleware(prv, routes.SetXRatedEnabledRoute(prv)))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:4568",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
