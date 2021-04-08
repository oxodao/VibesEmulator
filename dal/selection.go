package dal

import (
	"github.com/jmoiron/sqlx"
	"github.com/oxodao/vibes/models"
)

type Selection struct {
	DB  *sqlx.DB
	Dal *Dal
}

type DbSelection struct {
	SelectionID   int64  `db:"SELECTION_ID"`
	QuestionID    int64  `db:"QUESTION_ID"`
	UserAnswer    int64  `db:"USER_ANSWER"`
	UserText      string `db:"USER_COMMENT"`
	PartnerAnswer int64  `db:"PARTNER_ANSWER"`
	PartnerText   string `db:"PARTNER_COMMENT"`
}

func (s Selection) SelectionFromDbSelection(dbs DbSelection) (*models.Selection, error) {
	question, err := s.Dal.Questions.Find(dbs.QuestionID)
	if err != nil {
		return nil, err
	}

	realQuestion := &models.Selection{
		Question:      *question,
		Answers:       question.Answers,
		UserAnswer:    nil,
		PartnerAnswer: nil,
	}

	for _, a := range question.Answers {
		if a.ID == dbs.UserAnswer {
			realQuestion.UserAnswer = &a
		}

		if a.ID == dbs.UserAnswer {
			realQuestion.PartnerAnswer = &a
		}
	}

	return realQuestion, nil
}

func (s Selection) FindFromID(id int64) (*models.Selection, error) {
	row := s.DB.QueryRowx(`
		SELECT SELECTION_ID, QUESTION_ID, USER_ANSWER, USER_COMMENT, PARTNER_ANSWER, PARTNER_COMMENT
		FROM SELECTION
		WHERE SELECTION_ID = ?
`, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	dbSelection := DbSelection{}
	err := row.StructScan(&dbSelection)
	if err != nil {
		return nil, err
	}

	return s.SelectionFromDbSelection(dbSelection)
}

func (s Selection) New(phase *models.Phase) (*models.Selection, error) {
	q, err := s.Dal.Questions.FindRandom()
	if err != nil {
		return nil, err
	}

	_, err = s.DB.Exec(`
		INSERT INTO SELECTION(QUESTION_ID, PHASE_ID)
		VALUES (?, ?)
`, q.ID, phase.ID)

	if err != nil {
		return nil, err
	}

	return &models.Selection{
		Question: *q,
		Answers:  q.Answers,
	}, nil
}

func (s Selection) FindFromPhase(phase *models.Phase) error {
	rows, err := s.DB.Queryx(` 
		SELECT SELECTION_ID, QUESTION_ID, USER_ANSWER, USER_COMMENT, PARTNER_ANSWER, PARTNER_COMMENT
		FROM SELECTION
		WHERE PHASE_ID = ?
`, phase.ID)
	if err != nil {
		return err
	}

	selections := []models.Selection{}
	for rows.Next() {
		dbSelection := DbSelection{}
		err = rows.StructScan(&dbSelection)
		if err != nil {
			return err
		}

		selection, err := s.SelectionFromDbSelection(dbSelection)
		if err != nil {
			return err
		}

		selections = append(selections, *selection)
	}
	phase.Selections = selections

	return nil
}