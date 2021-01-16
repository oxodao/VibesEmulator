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
)

func Messenger(prv *services.Provider, r *mux.Router) {
	r.HandleFunc("/getMessages", middlewares.CheckUserMiddleware(prv, getMessagesRoute(prv)))
	r.HandleFunc("/addMessage", middlewares.CheckUserMiddleware(prv, sendMessageRoute(prv)))
}

type getMessageResponse struct {
	ButtonState string           `json:"buttonState"`
	FacebookID  string           `json:"facebookId"`
	Messages    []models.Message `json:"messages"`
}

func getMessagesRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		partner := r.URL.Query().Get("partnerId")
		partnerID, err := strconv.ParseUint(partner, 10, 64)
		if err != nil {
			// ??
		}

		msgs, err := dal.GetUsersMessage(prv, u, partnerID)
		if err != nil {
			// ??
		}

		rsp, _ := json.Marshal(getMessageResponse{
			//ButtonState: "WAIT",
			ButtonState: "PLAY",
			Messages:    msgs,
		})

		w.Write(rsp)

	}
}

func sendMessageRoute(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		user := r.URL.Query().Get("partnerId")
		text := r.URL.Query().Get("text")
		question := r.URL.Query().Get("question")

		if len(user) > 0 && len(text) > 0 {
			userID, err := strconv.ParseUint(user, 10, 64)

			if len(question) > 0 && question != "0" {
				// @TODO Find what's it used for
				questionID, err2 := strconv.Atoi(question)
				fmt.Println("Testt", questionID, err2)
			} else {
				if err != nil {
					fmt.Println("Messenger@SendMessageRoute: ", err)
					return
				}

				msg, err := dal.SendMessage(prv, u.ID, userID, text, "PLAIN", "", "")

				msgJSON, _ := json.Marshal(msg)
				if err != nil {
					return
				}

				w.Write(msgJSON)
			}
		}
	}
}
