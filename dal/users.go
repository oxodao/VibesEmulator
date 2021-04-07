package dal

import (
	"github.com/jmoiron/sqlx"
	"github.com/oxodao/vibes/models"
)

type User struct {
	DB  *sqlx.DB
	Dal *Dal
}

func (usr User) FindUserByToken(token string) (*models.User, error) {
	var user *models.User = &models.User{}
	row, err := usr.DB.Queryx("SELECT * FROM APP_USER WHERE LATEST_TOKEN = ?", token)
	if err != nil {
		return nil, err
	}

	row.Next()
	err = row.StructScan(user)
	row.Close()

	return user, err
}

func (usr User) FindUserByUsername(name string) (*models.User, error) {
	var user = models.User{}
	row := usr.DB.QueryRowx("SELECT * FROM APP_USER WHERE LOWER(USERNAME) = LOWER(?) LIMIT 1", name)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.StructScan(&user)
	return &user, err
}

func (usr User) SetLatestToken(uid uint64, token string) error {
	_, err := usr.DB.Exec("UPDATE APP_USER SET LATEST_TOKEN = ? WHERE ID = ?", token, uid)
	return err
}

func (usr User) RegisterUser(u *models.User) error {
	_, err := usr.DB.Exec(`
		INSERT INTO APP_USER(
			FIRSTNAME,
			PICTURE,
			AGE,
			GENDER,
			COUNTRY,
			LANG,
			PREMIUM,
			ADULT,
			AGE_FROM,
			AGE_TO,
			USERNAME,
			GENDER_WANTED,
			LAST_ACTION,
			PASSWORD,
			LATEST_TOKEN
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		u.FirstName,
		u.Picture,
		u.Age,
		u.Gender,
		u.Country,
		u.Language,
		u.IsPremium,
		u.IsAdult,
		u.AgeFrom,
		u.AgeTo,
		u.Username,
		u.GenderWanted,
		u.LastAction,
		u.Password,
		u.LatestToken)

	return err
}
