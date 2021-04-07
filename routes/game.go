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

		phase := models.NewPhase([]models.Selection{
			{
				Question: models.Question{
					ID:   0,
					Text: "Toto q0",
				},
				Answers: []models.Answer{
					{ID: 0, Text: "toto1", CommentText: nil},
					{ID: 1, Text: "toto2", CommentText: nil},
				},
				UserAnswer: nil,
				PartnerAnswer: nil,
			},
		})

		g := models.Game{
			Contact:             c,
			FinishedQuestionsID: []int64{},
			Phase0: &phase,
			Phase1: &phase,
			Phase2: &phase,
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
