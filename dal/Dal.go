package dal

import (
	"github.com/jmoiron/sqlx"
)

type Dal struct {
	Questions Questions
	Answers   Answers
	Selection Selection
	Phase     Phase
	User      User
	Contact   Contact
	Messenger Messenger
	Settings  Settings
}

func New(db *sqlx.DB) Dal {
	dal := &Dal{}

	dal.Questions = Questions{db, dal}
	dal.Answers = Answers{db, dal}
	dal.Selection = Selection{db, dal}
	dal.Phase = Phase{db, dal}
	dal.User = User{db, dal}
	dal.Contact = Contact{db, dal}
	dal.Messenger = Messenger{db, dal}
	dal.Settings = Settings{db, dal}

	return *dal
}
