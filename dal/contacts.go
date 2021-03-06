package dal

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	"github.com/oxodao/vibes/models"
)

type Contact struct {
	DB  *sqlx.DB
	Dal *Dal
}

// GenerateRandomContacts finds 5 random people to add to your suggestions.
// This should also saves them to get them five same until
func (c Contact) GenerateRandomContacts(uid uint64) ([]models.User, error) {
	rq := `
		SELECT *
		FROM APP_USER
		WHERE uid <> $1
		LIMIT 5
		ORDER BY random()
	`

	var users []models.User
	rows, err := c.DB.Queryx(rq, uid)
	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		u := models.User{}
		err = rows.StructScan(&u)

		if err != nil {
			break
		}

		users = append(users, u)
	}

	return users, err
}

func (c Contact) CreateOrFetchContactByName(user *models.User, username, webroot string) (models.Contact, error) {
	/**
		@TODO / Missing features:
			- Calculate the distance
			- Unread message counter
	**/
	rq := `
		SELECT IS_FRIENDLY, FRIEND_LEVEL, PROGRESS,
				INITIATOR.ID, INITIATOR.LAST_ACTION, INITIATOR.FIRSTNAME, INITIATOR.USERNAME, INITIATOR.GENDER, INITIATOR.COUNTRY, INITIATOR.AGE, INITIATOR.PICTURE, INITIATOR.LANG,
				FRIEND.ID, FRIEND.LAST_ACTION, FRIEND.FIRSTNAME, FRIEND.USERNAME, FRIEND.GENDER, FRIEND.COUNTRY, FRIEND.AGE, FRIEND.PICTURE, FRIEND.LANG
		FROM APP_CONTACTS c
				INNER JOIN APP_USER INITIATOR ON INITIATOR.ID = c.INITIATOR
				INNER JOIN APP_USER FRIEND ON FRIEND.ID = c.FRIEND
		WHERE (INITIATOR.ID = $1 OR FRIEND.ID = $1) AND (INITIATOR.ID = $2 OR FRIEND.ID = $2) 
	`

	rows, err := c.DB.Queryx(rq, user.ID, username)
	if err != nil {
		return models.Contact{}, err
	}

	if rows.Next() {
		contact := models.Contact{}
		rows.StructScan(&contact)

		contact.SetOtherUser(user)
		return contact, nil
	}

	rows.Close()

	user2, err := c.Dal.User.FindUserByUsername(username)
	if err != nil {
		return models.Contact{}, err
	}

	contact := models.Contact{
		Distance:   1,
		IsFriendly: true,
		Playable:   true,
		UserOne:    (*user).GetUserWithPictureURL(webroot),
		UserTwo:    user2.GetUserWithPictureURL(webroot),
		Turn:       1,
	}

	contact.SetOtherUser(user)

	rq = `INSERT INTO APP_CONTACTS (INITIATOR, FRIEND) VALUES (?, ?)`
	_, err = c.DB.Exec(rq, contact.UserOne.ID, contact.UserTwo.ID, true)

	if err != nil {
		fmt.Println(err)
		return models.Contact{}, err
	}

	return contact, nil
}

func (c Contact) GetContactsForUser(uid uint64, webroot string) ([]models.Contact, error) {
	contacts := []models.Contact{}

	rows, err := c.DB.Queryx(`SELECT 	IS_FRIENDLY, FRIEND_LEVEL, PROGRESS,
										i.ID as "INITIATOR.ID", i.LAST_ACTION as "INITIATOR.LAST_ACTION", i.FIRSTNAME as "INITIATOR.FIRSTNAME", i.USERNAME as "INITIATOR.USERNAME", i.GENDER as "INITIATOR.GENDER", i.COUNTRY as "INITIATOR.COUNTRY", i.AGE as "INITIATOR.AGE", i.PICTURE as "INITIATOR.PICTURE", i.LANG as "INITIATOR.LANG",
										f.ID as "FRIEND.ID", f.LAST_ACTION as "FRIEND.LAST_ACTION", f.FIRSTNAME as "FRIEND.FIRSTNAME", f.USERNAME as "FRIEND.USERNAME", f.GENDER as "FRIEND.GENDER", f.COUNTRY as "FRIEND.COUNTRY", f.AGE as "FRIEND.AGE", f.PICTURE as "FRIEND.PICTURE", f.LANG as "FRIEND.LANG"
								FROM APP_CONTACTS c
									LEFT JOIN APP_USER i ON i.ID = c.INITIATOR
									LEFT JOIN APP_USER f ON f.ID = c.FRIEND
								WHERE i.ID = $1 OR f.ID = $1`, uid)
	if err != nil {
		return contacts, err
	}

	contact := models.Contact{}
	for rows.Next() {
		rows.StructScan(&contact)
		contact.SetOtherUserID(uid)

		contact.User = contact.User.GetUserWithPictureURL(webroot)

		contacts = append(contacts, contact)
	}

	rows.Close()

	return contacts, nil
}

func (c Contact) GetContactByPartnerID(uid, uidFriend uint64, webroot string) (models.Contact, error) {
	contact := models.Contact{}

	row := c.DB.QueryRowx(`SELECT IS_FRIENDLY, FRIEND_LEVEL, PROGRESS,
									      i.ID as "INITIATOR.ID", i.LAST_ACTION as "INITIATOR.LAST_ACTION", i.FIRSTNAME as "INITIATOR.FIRSTNAME", i.USERNAME as "INITIATOR.USERNAME", i.GENDER as "INITIATOR.GENDER", i.COUNTRY as "INITIATOR.COUNTRY", i.AGE as "INITIATOR.AGE", i.PICTURE as "INITIATOR.PICTURE", i.LANG as "INITIATOR.LANG",
										  f.ID as "FRIEND.ID", f.LAST_ACTION as "FRIEND.LAST_ACTION", f.FIRSTNAME as "FRIEND.FIRSTNAME", f.USERNAME as "FRIEND.USERNAME", f.GENDER as "FRIEND.GENDER", f.COUNTRY as "FRIEND.COUNTRY", f.AGE as "FRIEND.AGE", f.PICTURE as "FRIEND.PICTURE", f.LANG as "FRIEND.LANG"
								FROM APP_CONTACTS c
									LEFT JOIN APP_USER i ON i.ID = c.INITIATOR
									LEFT JOIN APP_USER f ON f.ID = c.FRIEND
								WHERE (f.ID = $1 AND i.ID = $2) OR (f.ID = $2 AND i.ID = $1)`, uid, uidFriend)

	if row.Err() != nil {
		return contact, row.Err()
	}

	err := row.StructScan(&contact)
	if err != nil {
		return contact, err
	}

	contact.SetOtherUserID(uid)
	contact.User = contact.User.GetUserWithPictureURL(webroot)

	return contact, err
}
