package models

type Answer struct {
	ID          int64   `db:"ANSWER_ID" json:"id"`
	Text        string  `db:"ANSWER_TEXT" json:"text"`
	CommentText *string `db:"-" json:"commentText"`
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
