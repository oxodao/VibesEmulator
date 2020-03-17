package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/oxodao/vibes/middlewares"
	"github.com/oxodao/vibes/models"
	"github.com/oxodao/vibes/services"
	"github.com/oxodao/vibes/utils"
)

// SetPushTokenRoute blabla
func SetPushTokenRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Should return a user

		//		cookie, err := r.Cookie("PHPSESSID")

	}
}

// GetAllSettingsRoute blabla
func GetAllSettingsRoute(prv *services.Provider) http.HandlerFunc {
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

//SetAgeRangeRoute blabla
func SetAgeRangeRoute(prv *services.Provider) http.HandlerFunc {
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

		prv.DB.Model(&user).Update("age_from", ageFrom).Update("age_to", ageTo)

		userRsp, err := json.Marshal(user.GetUserWithPictureURL())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(userRsp)
	}
}

//SetAgeRoute blabla
func SetAgeRoute(prv *services.Provider) http.HandlerFunc {
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

		prv.DB.Model(&user).Update("age", age)

		userRsp, err := json.Marshal(user.GetUserWithPictureURL())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(userRsp)
	}
}

//SetFirstNameRoute blabla
func SetFirstNameRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		firstname := r.URL.Query().Get("firstName")

		user.FirstName = firstname

		prv.DB.Model(&user).Update("first_name", firstname)

		userRsp, err := json.Marshal(user.GetUserWithPictureURL())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(userRsp)
	}
}

//SetGenderRoute blabla
func SetGenderRoute(prv *services.Provider) http.HandlerFunc {
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

		prv.DB.Model(&user).Update("gender", gender)

		userRsp, err := json.Marshal(user.GetUserWithPictureURL())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(userRsp)
	}
}

//SetGenderWantedRoute blabla
func SetGenderWantedRoute(prv *services.Provider) http.HandlerFunc {
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

		prv.DB.Model(&user).Update("gender_wanted", gender)

		userRsp, err := json.Marshal(user.GetUserWithPictureURL())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(userRsp)
	}
}

// SetXRatedEnabledRoute blabla
func SetXRatedEnabledRoute(prv *services.Provider) http.HandlerFunc {
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
		prv.DB.Model(&user).Update("is_adult", user.IsAdult)

		resp, err := json.Marshal(user.GetUserWithPictureURL())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(resp)
	}
}

// SetGameLanguageRoute blabla
func SetGameLanguageRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		user.Language = r.URL.Query().Get("gameLanguage")

		prv.DB.Model(&user).Update("language", user.Language)

		resp, err := json.Marshal(user.GetUserWithPictureURL())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(resp)
	}
}

// SetPictureRoute uploads a picture
func SetPictureRoute(prv *services.Provider) http.HandlerFunc {
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

		prv.DB.Model(&user).Update("picture", rndName)

		uploadResponse, err := json.Marshal(user.GetUserWithPictureURL())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(uploadResponse)
	}
}
