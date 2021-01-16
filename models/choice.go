package models

type Choice struct {
	QuestionId uint64 `json:"questionId"`
	AnswerId uint64 `json:"answerId"`
}
