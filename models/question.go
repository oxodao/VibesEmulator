package models

type Question struct {
	ID      int64    `db:"QUESTION_ID" json:"id"`
	Text    string   `db:"QUESTION_TEXT" json:"text"`
	Answers []Answer `db:"-" json:"-"`
}

func (q Question) GetTableName() string {
	return "QUESTION"
}

func (q Question) GetTableCreationScript() string {
	return `
		CREATE TABLE QUESTION (
			QUESTION_ID INTEGER PRIMARY KEY AUTOINCREMENT,
			QUESTION_TEXT TEXT
		)
`
}
