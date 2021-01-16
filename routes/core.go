package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/oxodao/vibes/dal"
	"github.com/oxodao/vibes/middlewares"
	"github.com/oxodao/vibes/models"
	"github.com/oxodao/vibes/services"
	"github.com/oxodao/vibes/utils"
)

func Core(prv *services.Provider, r *mux.Router) {
	r.HandleFunc("/uploadPicture", uploadPictureRoute(prv))
	r.HandleFunc("/getContacts", middlewares.CheckUserMiddleware(prv, getContactsRoute(prv)))
	r.HandleFunc("/createContactWithUsername", middlewares.CheckUserMiddleware(prv, createContactWithUsernameRoute(prv)))
	r.HandleFunc("/createContactRandom", middlewares.CheckUserMiddleware(prv, createContactRandomRoute(prv)))
	//r.HandleFunc("/getPotentialContacts", middlewares.CheckUserMiddleware(prv, routes.GetPotentialContactsRoute(prv)))
}

func uploadPictureRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// @TODO: Store in DB, if not affected to a user after 5 minutes => Remove it
		// If affected => remove from DB
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

		uploadResponse, err := json.Marshal(struct {
			Picture string `json:"picture"`
		}{
			Picture: rndName,
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(uploadResponse)
	}
}

// /core/getContacts?latitude=49.24324324324324&longitude=4.057061128407301&country=FR&language=fr&os=android&version=12

/*
   private List<String> availableGameLanguages;
   private List<Contact> contacts;
   private List<Contact> potentialContacts;
   private boolean randomGameSearchOngoing;
   private String userGameLanguage;
*/

// getContactsRoute returns the friends + a tinder-like list of people
func getContactsRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		// @TODO: Store them so that we don't have random each time and only refill the list when the user calls this route
		randUsers, err := dal.GenerateRandomContacts(prv, u.ID)

		if err != nil {
			//??
		}

		var randContacts = make([]models.Contact, len(randUsers))
		for i := 0; i < len(randUsers); i++ {
			randContacts[i] = models.Contact{
				User: randUsers[i].GetUserWithPictureURL(prv.Config.WebrootURL),
			}
		}

		myContacts, err := dal.GetContactsForUser(prv, u.ID)
		if err != nil {
			// ??
		}

		getContactsResponse, err := json.Marshal(struct {
			AvailableGameLanguages  []string         `json:"availableGameLanguages"`
			Contacts                []models.Contact `json:"contacts"`
			PotentialContacts       []models.Contact `json:"potentialContacts"`
			RandomGameSearchOngoing bool             `json:"randomGameSearchOngoing"`
			UserGameLanguage        string           `json:"userGameLanguage"`
		}{
			AvailableGameLanguages:  []string{"fr"},
			UserGameLanguage:        "fr",
			RandomGameSearchOngoing: false,
			Contacts:                myContacts,
			PotentialContacts:       randContacts,
		})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(getContactsResponse)
	}
}

func createContactRandomRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		// Temporary
		// Should be only where "randomSearchOngoing" I think
		var rndContact models.Contact

		resp, err := json.Marshal(rndContact)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(resp)
	}
}

func createContactWithUsernameRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		otherUser := r.URL.Query().Get("username")
		rndContact, err := dal.CreateOrFetchContactByName(prv, u, otherUser)

		resp, err := json.Marshal(rndContact)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(resp)
	}
}
