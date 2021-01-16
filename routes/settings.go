package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/oxodao/vibes/dal"
	"github.com/oxodao/vibes/middlewares"
	"github.com/oxodao/vibes/models"
	"github.com/oxodao/vibes/services"
	"github.com/oxodao/vibes/utils"
)

func Settings(prv *services.Provider, r *mux.Router) {
	r.HandleFunc("/getAll", middlewares.CheckUserMiddleware(prv, getAllSettingsRoute(prv)))
	r.HandleFunc("/setAge", middlewares.CheckUserMiddleware(prv, setAgeRoute(prv)))
	r.HandleFunc("/setAgeRange", middlewares.CheckUserMiddleware(prv, setAgeRangeRoute(prv)))
	r.HandleFunc("/setFirstName", middlewares.CheckUserMiddleware(prv, setFirstNameRoute(prv)))
	r.HandleFunc("/setGameLanguage", middlewares.CheckUserMiddleware(prv, setGameLanguageRoute(prv)))
	r.HandleFunc("/setGender", middlewares.CheckUserMiddleware(prv, setGenderRoute(prv)))
	r.HandleFunc("/setGenderWanted", middlewares.CheckUserMiddleware(prv, setGenderWantedRoute(prv)))
	r.HandleFunc("/setPicture", middlewares.CheckUserMiddleware(prv, setPictureRoute(prv)))
	r.HandleFunc("/setPushToken", middlewares.CheckUserMiddleware(prv, setPushTokenRoute(prv)))
	r.HandleFunc("/setXRatedEnabled", middlewares.CheckUserMiddleware(prv, setXRatedEnabledRoute(prv)))
}

func setPushTokenRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Should return a user

		//		cookie, err := r.Cookie("PHPSESSID")

	}
}

func getAllSettingsRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		resp, err := json.Marshal(struct {
			AvailableLanguage []string    `json:"availableGameLanguages"`
			User              models.User `json:"user"`
		}{
			AvailableLanguage: []string{"fr"},
			User:              user.GetUserWithPictureURL(),
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(resp)
	}
}

func setAgeRangeRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		ageFrom, err1 := strconv.Atoi(r.URL.Query().Get("ageFrom"))
		ageTo, err2 := strconv.Atoi(r.URL.Query().Get("ageTo"))
		if err1 != nil || err2 != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user.AgeFrom = ageFrom
		user.AgeTo = ageTo
		dal.UpdateAge(prv, user.ID, -1, ageFrom, ageTo)

		userRsp, err := json.Marshal(user.GetUserWithPictureURL())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(userRsp)
	}
}

func setAgeRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		age, err := strconv.Atoi(r.URL.Query().Get("age"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user.Age = age
		dal.UpdateAge(prv, user.ID, age, -1, -1)

		userRsp, err := json.Marshal(user.GetUserWithPictureURL())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(userRsp)
	}
}

func setFirstNameRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		firstname := r.URL.Query().Get("firstName")

		user.FirstName = firstname
		dal.UpdateFirstname(prv, user.ID, firstname)

		userRsp, err := json.Marshal(user.GetUserWithPictureURL())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(userRsp)
	}
}

func setGenderRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		gender, err := strconv.Atoi(r.URL.Query().Get("gender"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user.Gender = gender
		dal.UpdateGender(prv, user.ID, gender, -1)

		userRsp, err := json.Marshal(user.GetUserWithPictureURL())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(userRsp)
	}
}

func setGenderWantedRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		gender, err := strconv.Atoi(r.URL.Query().Get("genderWanted"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user.GenderWanted = gender

		dal.UpdateGender(prv, user.ID, -1, gender)

		userRsp, err := json.Marshal(user.GetUserWithPictureURL())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(userRsp)
	}
}

func setXRatedEnabledRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		enabled := r.URL.Query().Get("xRatedEnabled")

		if enabled == "0" {
			user.IsAdult = false
		} else {
			user.IsAdult = true
		}
		dal.UpdateAdult(prv, user.ID, user.IsAdult)

		resp, err := json.Marshal(user.GetUserWithPictureURL())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(resp)
	}
}

func setGameLanguageRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		user.Language = r.URL.Query().Get("gameLanguage")
		dal.UpdateLanguage(prv, user.ID, user.Language)

		resp, err := json.Marshal(user.GetUserWithPictureURL())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(resp)
	}
}

func setPictureRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		r.ParseMultipartForm(32 << 20)

		file, _, err := r.FormFile("picture")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer file.Close()

		rndName, err := utils.SetPicture(prv, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		dal.UpdatePicture(prv, user.ID, rndName)

		uploadResponse, err := json.Marshal(user.GetUserWithPictureURL())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(uploadResponse)
	}
}
