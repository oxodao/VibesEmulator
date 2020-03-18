package dal

import (
	"fmt"

	"github.com/oxodao/vibes/models"
	"github.com/oxodao/vibes/services"
)

// GenerateRandomContacts finds 5 random people to add to your suggestions.
// This should also saves them to get them five same until
func GenerateRandomContacts(prv *services.Provider, uid uint) ([]models.User, error) {
	rq := `
		SELECT *
		FROM APP_USER
		WHERE uid <> $1
		LIMIT 5
		ORDER BY random()
	`

	var users []models.User
	rows, err := prv.DB.Queryx(rq, uid)
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

// CreateOrFetchContactByName blabla
func CreateOrFetchContactByName(prv *services.Provider, user *models.User, username string) (models.Contact, error) {
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

	rows, err := prv.DB.Queryx(rq, user.ID, username)
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

	rq = `SELECT * FROM APP_USER WHERE LOWER(USERNAME) = LOWER(?) `

	user2 := models.User{}
	err = prv.DB.Get(&user2, rq, username)
	if err != nil {
		return models.Contact{}, err
	}

	contact := models.Contact{
		Distance:   1,
		IsFriendly: true,
		Playable:   true,
		UserOne:    (*user).GetUserWithPictureURL(),
		UserTwo:    user2.GetUserWithPictureURL(),
		Turn:       1,
	}

	contact.SetOtherUser(user)

	rq = `INSERT INTO APP_CONTACTS (INITIATOR, FRIEND) VALUES (?, ?)`
	_, err = prv.DB.Exec(rq, contact.UserOne.ID, contact.UserTwo.ID, true)

	if err != nil {
		fmt.Println(err)
		return models.Contact{}, err
	}

	return contact, nil
}

// GetContactsForUser blabla
func GetContactsForUser(prv *services.Provider, uid uint) ([]models.Contact, error) {
	contacts := []models.Contact{}

	rows, err := prv.DB.Queryx(`SELECT 	IS_FRIENDLY, FRIEND_LEVEL, PROGRESS,
										INITIATOR.ID, INITIATOR.LAST_ACTION, INITIATOR.FIRSTNAME, INITIATOR.USERNAME, INITIATOR.GENDER, INITIATOR.COUNTRY, INITIATOR.AGE, INITIATOR.PICTURE, INITIATOR.LANG,
										FRIEND.ID, FRIEND.LAST_ACTION, FRIEND.FIRSTNAME, FRIEND.USERNAME, FRIEND.GENDER, FRIEND.COUNTRY, FRIEND.AGE, FRIEND.PICTURE, FRIEND.LANG
								FROM APP_CONTACTS c
									LEFT JOIN APP_USER INITIATOR ON INITIATOR.ID = c.INITIATOR
									LEFT JOIN APP_USER FRIEND ON FRIEND.ID = c.FRIEND
								WHERE INITIATOR.ID = $1 OR FRIEND.ID = $1`, uid)
	if err != nil {
		return contacts, err
	}

	for rows.Next() {
		var contact models.Contact
		rows.StructScan(&contact)

		contact.SetOtherUserID(uid)
		fmt.Printf("\t- %v / %v\n", contact.UserOne.Username, contact.UserTwo.Username)

		contacts = append(contacts, contact)
	}

	rows.Close()

	return contacts, nil
}
