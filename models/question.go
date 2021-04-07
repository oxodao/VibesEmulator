package models

type Question struct {
	ID      uint64   `json:"id"`
	Text    string   `json:"text"`
	Answers []Answer `json:"-"`
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