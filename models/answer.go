package models

type Answer struct {
	ID          uint64 `json:"id"`
	Text        string `json:"text"`
	CommentText *string `json:"commentText"`
}

func (a Answer) GetTableName() string {
	return "ANSWER"
}

func (a Answer) GetTableCreationScript() string {
	return `
		CREATE TABLE ANSWER (
			ANSWER_ID INTEGER PRIMARY KEY AUTOINCREMENT,
			ANSWER_TEXT TEXT,
			QUESTION_ID INTEGER REFERENCES QUESTION(QUESTION_ID)
		)
`
}
