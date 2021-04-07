package models

// Contact represent a friend
type Contact struct {
	ID                 int  `json:"id"`
	Distance           int  `json:"distance"`
	IsFriendly         bool `json:"isFriendly" db:"IS_FRIENDLY"`
	Level              int  `json:"level" db:"FRIEND_LEVEL"`
	Playable           bool `json:"playable" db:"-"`
	Progress           int  `json:"progress" db:"PROGRESS"`
	LastAction         uint `db:"LAST_ACTION"`
	UnreadMessageCount int  `json:"unreadMessageCount" db:"-"`
	User               User `json:"user" db:"-"`
	Turn               int  `json:"-" db:"TURN"`
	UserOne            User `json:"-" db:"INITIATOR"`
	UserTwo            User `json:"-" db:"FRIEND"`
}

// SetOtherUser sets the user to the contact's user
func (c *Contact) SetOtherUser(currUser *User) {
	c.SetOtherUserID(currUser.ID)
}

// SetOtherUserID sets the user to the contact's user
func (c *Contact) SetOtherUserID(uid uint64) {
	if c.UserOne.ID == uid {
		c.User = c.UserTwo
		c.Playable = c.Turn == 1
	} else {
		c.User = c.UserOne
		c.Playable = c.Turn == 2
	}

}


func (c Contact) GetTableName() string {
	return "APP_CONTACTS"
}

func (c Contact) GetTableCreationScript() string {
	return `
		CREATE TABLE APP_CONTACTS (
			INITIATOR INTEGER,
			FRIEND INTEGER,
			IS_FRIENDLY BOOL DEFAULT 1,
			FRIEND_LEVEL INTEGER DEFAULT 0,
			TURN INTEGER DEFAULT 1,
			PROGRESS INTEGER DEFAULT 0,
			PRIMARY KEY (INITIATOR, FRIEND),
			FOREIGN KEY(INITIATOR) REFERENCES APP_USER(ID),
			FOREIGN KEY(FRIEND) REFERENCES APP_USER(ID)
		);
`
}