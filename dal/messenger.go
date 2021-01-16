package dal

import (
	"fmt"

	"github.com/oxodao/vibes/models"
	"github.com/oxodao/vibes/services"
)

/**
 * @TODO: Something's wrong with this
 * The sender sees the sent message as received and vice versa
 */

func GetUsersMessage(prv *services.Provider, u *models.User, p uint64) ([]models.Message, error) {
	rows, err := prv.DB.Queryx(`
			SELECT m.ID, m.CREATED_AT, m.TYPE, m.CONTENT, m.SENDER_ANSWER_TEXT, m.RECEIVER_ANSWER_TEXT,
				   su.ID as "SENDER", ru.ID as "RECEIVER"
			FROM APP_MESSENGER m
				INNER JOIN APP_USER su ON su.ID = m.SENDER
				INNER JOIN APP_USER ru ON ru.ID = m.RECEIVER
			WHERE
				(su.ID = $1 AND ru.ID = $2)
				OR
				(su.ID = $2 AND ru.ID = $1)
			ORDER BY m.CREATED_AT DESC 
	`, u.ID, p)

	if err != nil {
		fmt.Println("Messenger@GetUsersMessage -> ", err)
		return []models.Message{}, err
	}

	messages := []models.Message{}
	for rows.Next() {
		var msg models.Message
		err = rows.StructScan(&msg)
		msg.SetOtherUser(u)

		messages = append(messages, msg)
	}

	rows.Close()
	return messages, nil
}

func SendMessage(prv *services.Provider, from uint64, to uint64, message string, msgType string, senderAnswerText string, receiverAnswerText string) (models.Message, error) {
	inserted, err := prv.DB.Exec(`INSERT INTO 	APP_MESSENGER (
										SENDER,
										RECEIVER,
										CONTENT,
										TYPE,
										SENDER_ANSWER_TEXT,
										RECEIVER_ANSWER_TEXT
		) VALUES (?, ?, ?, ?, ?, ?)`, from, to, message, msgType, senderAnswerText, receiverAnswerText)

	if err != nil {
		return models.Message{}, err
	}

	id, err := inserted.LastInsertId()
	if err != nil {
		return models.Message{}, err
	}

	row := prv.DB.QueryRowx("SELECT * FROM APP_MESSENGER WHERE ID = ?", id)

	var msg models.Message
	row.StructScan(&msg)

	return msg, err
}
