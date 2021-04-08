package models

type Selection struct {
	Question      Question `json:"question"`
	Answers       []Answer `json:"answers"`
	UserAnswer    *Answer  `json:"userAnswer"`
	PartnerAnswer *Answer  `json:"partnerAnswer"`
}

func (s Selection) GetTableName() string {
	return "SELECTION"
}

func (s Selection) GetTableCreationScript() string {
	return `
		CREATE TABLE SELECTION (
			SELECTION_ID INTEGER PRIMARY KEY AUTOINCREMENT,
			PHASE_ID     INTEGER REFERENCES PHASE(PHASE_ID),
			QUESTION_ID  INTEGER REFERENCES QUESTION(QUESTION_ID),
			USER_ANSWER  INTEGER REFERENCES ANSWER(ANSWER_ID),
			USER_COMMENT TEXT,
			PARTNER_ANSWER  INTEGER REFERENCES ANSWER(ANSWER_ID),
			PARTNER_COMMENT TEXT
);
`
}