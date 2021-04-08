package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/oxodao/vibes/middlewares"
	"github.com/oxodao/vibes/models"
	"github.com/oxodao/vibes/services"
	"net/http"
	"strconv"
)

func Game(prv *services.Provider, r *mux.Router) {
	r.HandleFunc("/getData", middlewares.CheckUserMiddleware(prv, getData(prv)))
}

func getData(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, ok := r.Context().Value(middlewares.UserContext).(*models.User)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Can't cast the user!")
			return
		}

		partnerIDStr := r.URL.Query().Get("partnerId")
		partnerID, err := strconv.ParseUint(partnerIDStr, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println("Partner: ", partnerID)

		c, err := prv.Dal.Contact.GetContactByPartnerID(u.ID, partnerID, prv.Config.WebrootURL)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		phases, err := prv.Dal.Phase.FindLatestPhases(&c)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)

			return
		}

		g := models.Game{
			Contact:             c,
			FinishedQuestionsID: []int64{},
			Phase0: phases[2],
			Phase1: phases[1],
			Phase2: phases[0],
			ProgressCalculation: models.ProgressCalculation{
				A:    0,
				B:    0,
				Loss: 0,
			},
			UserChoices: []models.Choice{
				{QuestionId: 0, AnswerId: 0},
			},
		}

		gBytes, _ := json.Marshal(&g)
		w.Write(gBytes)
	}
}
