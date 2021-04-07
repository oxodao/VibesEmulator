package models

// Message represents a message. Wow.
type Message struct {
	ID                uint   `json:"id" db:"ID"`
	Date              uint64 `json:"time" db:"CREATED_AT"`
	SentByUserOne     bool   `json:"isFromUser" db:"-"`
	Type              string `json:"type" db:"TYPE"` // Can be PLAIN, QUESTION or COMMENT
	Text              string `json:"text" db:"CONTENT"`
	UserAnswerText    string `json:"userAnswerText" db:"-"`
	PartnerAnswerText string `json:"partnerAnswerText" db:"-"`

	Sender       uint64 `json:"-" db:"SENDER"`
	Receiver     uint64 `json:"-" db:"RECEIVER"`
	SenderText   string `json:"-" db:"SENDER_ANSWER_TEXT"`
	ReceiverText string `json:"-" db:"RECEIVER_ANSWER_TEXT"`
}

// SetOtherUser blabla
func (m *Message) SetOtherUser(mainUser *User) {
	if m.Sender == mainUser.ID {
		m.UserAnswerText = m.SenderText
		m.PartnerAnswerText = m.ReceiverText
	} else {
		m.SentByUserOne = true
		m.UserAnswerText = m.ReceiverText
		m.PartnerAnswerText = m.SenderText
	}
}


func (m Message) GetTableName() string {
	return "APP_MESSENGER"
}

func (m Message) GetTableCreationScript() string {
	return `
		CREATE TABLE APP_MESSENGER (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			SENDER INTEGER,
			RECEIVER INTEGER,
			CONTENT VARCHAR(255),
			TYPE VARCHAR(10),
			SENDER_ANSWER_TEXT VARCHAR(255),
			RECEIVER_ANSWER_TEXT VARCHAR(255),
			CREATED_AT INTEGER DEFAULT (cast(strftime('%s','now') as int)),
			FOREIGN KEY(SENDER) REFERENCES APP_USER(ID),
			FOREIGN KEY(RECEIVER) REFERENCES APP_USER(ID)
		);
`
}