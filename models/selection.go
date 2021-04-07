package models

type Selection struct {
	Question      Question `json:"question"`
	Answers       []Answer `json:"answers"`
	UserAnswer    *Answer  `json:"userAnswer"`
	PartnerAnswer *Answer  `json:"partnerAnswer"`
}
