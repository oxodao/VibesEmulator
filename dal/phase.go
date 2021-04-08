package dal

import (
	"github.com/jmoiron/sqlx"
	"github.com/oxodao/vibes/models"
)

type Phase struct {
	DB  *sqlx.DB
	Dal *Dal
}

func (p Phase) New(contact *models.Contact) (*models.Phase, error) {
	phase := models.Phase{
		ContactID: contact.ID,
	}

	resp, err := p.DB.Exec(`INSERT INTO PHASE (CONTACT_ID) VALUES (?)`, phase.ContactID)
	if err != nil {
		return nil, err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return nil, err
	}

	phase.ID = id

	selection1, err := p.Dal.Selection.New(&phase)
	if err != nil {
		return nil, err
	}

	selection2, err := p.Dal.Selection.New(&phase)
	if err != nil {
		return nil, err
	}

	phase.Selections = []models.Selection{
		*selection1,
		*selection2,
	}

	return &phase, nil
}

func (p Phase) FindLatestPhases(contact *models.Contact) ([]*models.Phase, error) {
	phases := []*models.Phase{}

	rows, err := p.DB.Queryx(`
		SELECT PHASE_ID, CONTACT_ID
		FROM PHASE
		WHERE CONTACT_ID = ?
		ORDER BY PHASE_ID DESC
		LIMIT 3
	`, contact.ID)
	if err != nil {
		return phases, err
	}

	for rows.Next() {
		currPhase := models.Phase{}
		err := rows.StructScan(&currPhase)
		if err != nil {
			return phases, err
		}

		err = p.Dal.Selection.FindFromPhase(&currPhase)
		if err != nil {
			return phases, err
		}

		phases = append(phases, &currPhase)
	}

	if len(phases) == 0 {
		for i := 0; i < 3; i++ {
			phase, err := p.New(contact)
			if err != nil {
				return phases, err
			}

			phases = append(phases, phase)
		}
	}

	for i := 3 - len(phases); i > 0; i-- {
		phases = append(phases, nil)
	}

	return phases, nil
}
