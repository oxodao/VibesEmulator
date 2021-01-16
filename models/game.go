package models

type Game struct {
	Contact             Contact             `json:"contact"`
	FinishedQuestionsID []int64             `json:"finishedQuestionId"`
	Phase0              *Phase               `json:"phase0"`
	Phase1              *Phase               `json:"phase1"`
	Phase2              *Phase               `json:"phase2"`
	ProgressCalculation ProgressCalculation `json:"progressCalculation"`
	UserChoices         []Choice            `json:"userChoices"`
}
