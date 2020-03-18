package dal

import (
	"github.com/oxodao/vibes/models"
	"github.com/oxodao/vibes/services"
)

// FindUserByToken blabla
func FindUserByToken(prv *services.Provider, token string) (*models.User, error) {
	var user *models.User = &models.User{}
	row, err := prv.DB.Queryx("SELECT * FROM APP_USER WHERE LATEST_TOKEN = ?", token)
	if err != nil {
		return nil, err
	}

	row.Next()
	err = row.StructScan(user)
	row.Close()

	return user, err
}

// FindUserByUsername blabla
func FindUserByUsername(prv *services.Provider, name string) (*models.User, error) {
	var user *models.User = &models.User{}
	row, err := prv.DB.Queryx("SELECT * FROM APP_USER WHERE LOWER(USERNAME) = LOWER(?)", name)
	if err != nil {
		return nil, err
	}

	row.Next()
	err = row.StructScan(user)
	row.Close()

	return user, err
}

// SetLatestToken blabla
func SetLatestToken(prv *services.Provider, uid uint, token string) error {
	_, err := prv.DB.Exec("UPDATE APP_USER SET LATEST_TOKEN = ? WHERE ID = ?", token, uid)
	return err
}

// RegisterUser blabla
func RegisterUser(prv *services.Provider, u *models.User) error {
	_, err := prv.DB.Exec(`
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
