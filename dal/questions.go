package dal

import (
	"github.com/jmoiron/sqlx"
	"github.com/oxodao/vibes/models"
)

type Questions struct {
	DB  *sqlx.DB
	Dal *Dal
}

func (q Questions) InsertQuestion(question *models.Question) error {
	resp, err := q.DB.Exec(`INSERT INTO QUESTION (QUESTION_TEXT) VALUES (?)`, question.Text)
	if err != nil {
		return err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return err
	}

	question.ID = id

	for _, a := range question.Answers {
		err = q.Dal.Answers.Insert(&a, id)
	}

	return nil
}

func (q Questions) Find(id int64) (*models.Question, error) {
	row := q.DB.QueryRowx(`SELECT QUESTION_ID, QUESTION_TEXT FROM QUESTION WHERE QUESTION_ID = ?`, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	question := models.Question{}
	err := row.StructScan(&question)
	if err != nil {
		return nil, err
	}

	question.Answers, err = q.Dal.Answers.FindForQuestion(question.ID)
	return &question, err
}

func (q Questions) FindRandom() (*models.Question, error) {
	row := q.DB.QueryRowx(`SELECT QUESTION_ID, QUESTION_TEXT FROM QUESTION ORDER BY RANDOM() LIMIT 1`)
	if row.Err() != nil {
		return nil, row.Err()
	}

	question := models.Question{}
	err := row.StructScan(&question)
	if err != nil {
		return nil, err
	}

	question.Answers, err = q.Dal.Answers.FindForQuestion(question.ID)
	return &question, err
}