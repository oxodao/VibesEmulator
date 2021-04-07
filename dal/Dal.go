package dal

import (
	"github.com/jmoiron/sqlx"
)

type Dal struct {
	Questions Questions
	User      User
	Contact   Contact
	Messenger Messenger
	Settings  Settings
}

func New(db *sqlx.DB) Dal {
	return Dal{
		Questions: Questions{db},
		User:      User{db},
		Contact:   Contact{db},
		Messenger: Messenger{db},
		Settings:  Settings{db},
	}
}
