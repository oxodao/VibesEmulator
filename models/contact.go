package models

// Contact represent a friend
type Contact struct {
	Distance           int  `json:"distance"`
	IsFriendly         bool `json:"isFriendly"`
	Level              int  `json:"level"`
	Playable           bool `json:"playable"`
	Progress           int  `json:"progress"`
	UnreadMessageCount int  `json:"unreadMessageCount"`
	User               User `json:"user" gorm:"-"`
	UserOne            User `json:"-"`
	UserTwo            User `json:"-"`
}

// SetOtherUser sets the user to the contact's user
func (c *Contact) SetOtherUser(currUser *User) {
	if c.UserOne.ID == currUser.ID {
		c.User = c.UserTwo
	} else {
		c.User = c.UserOne
	}
}
