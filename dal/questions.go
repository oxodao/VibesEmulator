package dal

import (
	"github.com/jmoiron/sqlx"
	"github.com/oxodao/vibes/models"
)

type Questions struct {
	DB *sqlx.DB
}

func (q Questions) InsertQuestion(question models.Question) error {
	resp, err := q.DB.Exec(`INSERT INTO QUESTION (QUESTION_TEXT) VALUES (?)`, question.Text)
	if err != nil {
		return err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return err
	}

	for _, a := range question.Answers {
		_, err := q.DB.Exec(`INSERT INTO ANSWER (ANSWER_TEXT, QUESTION_ID) VALUES (?, ?)`, a.Text, id)
		if err != nil {
			return err
		}
	}

	return nil
}