package dal

import (
	"github.com/jmoiron/sqlx"
	"github.com/oxodao/vibes/models"
)

type Answers struct {
	DB  *sqlx.DB
	Dal *Dal
}

func (a Answers) Insert(answer *models.Answer, questionId int64) error {
	resp, err := a.DB.Exec(`INSERT INTO ANSWER (ANSWER_TEXT, QUESTION_ID) VALUES (?, ?)`, answer.Text, questionId)
	if err != nil {
		return err
	}

	idAnswer, err := resp.LastInsertId()
	if err != nil {
		return err
	}

	answer.ID = idAnswer

	return nil
}

func (a Answers) Find(id int64) (*models.Answer, error) {
	return nil, nil
}

func (a Answers) FindForQuestion(questionId int64) ([]models.Answer, error) {
	rows, err := a.DB.Queryx(`
		SELECT ANSWER_ID, ANSWER_TEXT
		FROM ANSWER
		WHERE QUESTION_ID = ?
	`, questionId)

	if err != nil {
		return nil, err
	}

	answers := []models.Answer{}
	for rows.Next() {
		answer := models.Answer{}
		err := rows.StructScan(&answer)
		if err != nil {
			return nil, err
		}

		answers = append(answers, answer)
	}

	return answers, nil
}