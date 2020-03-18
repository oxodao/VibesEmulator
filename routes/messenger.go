package routes

import (
	"encoding/json"
	"net/http"

	"github.com/oxodao/vibes/services"
)

// GetMessages blalba
func GetMessages(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type Response struct {
			ButtonState string   `json:"buttonState"`
			FacebookID  string   `json:"facebookId"`
			Messages    []string `json:"messages"`
		}

		rsp, _ := json.Marshal(Response{
			ButtonState: "WAIT",
			Messages:    []string{},
		})

		w.Write(rsp)

	}
}
