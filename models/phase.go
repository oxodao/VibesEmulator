package models

type Phase struct {
	ID         int64       `db:"PHASE_ID" json:"-"`
	ContactID  int64       `db:"CONTACT_ID" json:"-"`
	Selections []Selection `db:"-" json:"selections"`
}

func (p Phase) GetTableName() string {
	return "PHASE"
}

func (p Phase) GetTableCreationScript() string {
	return `
		CREATE TABLE PHASE (
			PHASE_ID   INTEGER PRIMARY KEY AUTOINCREMENT,
			CONTACT_ID INTEGER REFERENCES CONTACT(CONTACT_ID)
		)
`
}