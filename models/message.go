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

	Sender       uint   `json:"-" db:"SENDER"`
	Receiver     uint   `json:"-" db:"RECEIVER"`
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