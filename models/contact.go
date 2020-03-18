package models

// Contact represent a friend
type Contact struct {
	Distance           int  `json:"distance"`
	IsFriendly         bool `json:"isFriendly" db:"IS_FRIENDLY"`
	Level              int  `json:"level" db:"FRIEND_LEVEL"`
	Playable           bool `json:"playable" db:"-"`
	Progress           int  `json:"progress" db:"PROGRESS"`
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
func (c *Contact) SetOtherUserID(uid uint) {
	if c.UserOne.ID == uid {
		c.User = c.UserTwo
		c.Playable = c.Turn == 1
	} else {
		c.User = c.UserOne
		c.Playable = c.Turn == 2
	}

}
