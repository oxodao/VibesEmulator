package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"

	"github.com/oxodao/vibes/middlewares"
	"github.com/oxodao/vibes/models"
	"github.com/oxodao/vibes/services"
)

func Auth(prv *services.Provider, r *mux.Router) {
	r.HandleFunc("/register", registerRoute(prv))
	r.HandleFunc("/login", loginRoute(prv))
	r.HandleFunc("/logout", logoutRoute(prv))
}

func registerRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("New user !")
		q := r.URL.Query()

		age := q.Get("age")
		ageInt, err := strconv.Atoi(age)
		if err != nil {
			ageInt = 20
		}

		gender := q.Get("gender")
		genderInt, err := strconv.Atoi(gender)
		if err != nil {
			genderInt = 1
		}

		genderWanted := 0
		if genderInt == 0 {
			genderWanted = 1
		}

		pwd := prv.GenerateUID(32)
		pwdHashed, err := prv.HashPassword(pwd)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		/** @TODO: Replace LatestToken with a many2one to allow multiple devices to be logged in **/
		/**
			Must not update the Picture with the full URL, this is the one received after uploading it directly
			CF. de.lotumapps.vibes.ui.activity.RegisterActivity.onRegister(View arg6)
		**/
		u := models.User{
			FirstName:    q.Get("firstName"),
			Picture:      q.Get("picture"),
			Age:          ageInt,
			Gender:       genderInt,
			Country:      q.Get("country"),
			Language:     q.Get("language"),
			IsPremium:    false,
			IsAdult:      ageInt > 18,
			AgeFrom:      ageInt - 2,
			AgeTo:        ageInt + 2,
			Username:     q.Get("firstName") + "_" + prv.GenerateUID(5),
			GenderWanted: genderWanted,
			LastAction:   time.Now().Unix(),
			Password:     pwdHashed,
			LatestToken:  prv.GenerateUID(20),
		}

		prv.Dal.User.RegisterUser(&u)

		http.SetCookie(w, &http.Cookie{
			Name:    "PHPSESSID",
			Value:   u.LatestToken,
			Expires: time.Now().Add(365 * 24 * time.Hour),
			Path:    "/",
		})

		registerResponse, err := json.Marshal(struct {
			User     models.User `json:"user"`
			Password string      `json:"password"`
		}{
			User:     u.GetUserWithPictureURL(prv.Config.WebrootURL),
			Password: pwd,
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(registerResponse)
	}
}

func loginRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		password := r.URL.Query().Get("password")

		u, err := prv.Dal.User.FindUserByUsername(username)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		pwdHashed, err := prv.HashPassword(password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if pwdHashed != u.Password {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		userJSON, err := json.Marshal(u)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(userJSON)
	}
}

func logoutRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		u, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		// The game says the data will be wiped from the server but I think we can find
		// a way around that to let the user log back in again
		prv.Dal.User.SetLatestToken(u.ID, "")

	}
}
