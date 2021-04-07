package dal

import (
	"github.com/jmoiron/sqlx"
)

/**
@TODO Handle errors
**/

type Settings struct {
	DB *sqlx.DB
}

func (s Settings) UpdateAge(uid uint64, age int, ageFrom int, ageTo int) {
	s.DB.Exec(`
		UPDATE APP_USER
		SET
			AGE = CASE WHEN $1 <> -1 THEN $1 ELSE AGE END,
			AGE_FROM = CASE WHEN $2 <> -1 THEN $2 ELSE AGE END,
			AGE_TO = CASE WHEN $3 <> -1 THEN $3 ELSE AGE END
		WHERE ID = $4
	`, age, ageFrom, ageTo, uid)
}

func (s Settings) UpdateFirstname(uid uint64, name string) {
	s.DB.Exec("UPDATE APP_USER SET FIRSTNAME = ? WHERE ID = ?", name, uid)
}

func (s Settings) UpdateGender(uid uint64, gender int, genderWanted int) {
	s.DB.Exec(`
		UPDATE APP_USER
		SET
			GENDER = CASE WHEN $1 <> -1 THEN $1 ELSE GENDER END,
			GENDER_WANTED = CASE WHEN $2 <> -1 THEN $2 ELSE GENDER_WANTED END
		WHERE ID = $3
	`, gender, genderWanted, uid)
}

func (s Settings) UpdateAdult(uid uint64, adult bool) {
	s.DB.Exec("UPDATE APP_USER SET ADULT = ? WHERE ID = ?", adult, uid)
}

func (s Settings) UpdateLanguage(uid uint64, lang string) {
	s.DB.Exec("UPDATE APP_USER SET LANG = ? WHERE ID = ?", lang, uid)
}

func (s Settings) UpdatePicture(uid uint64, pict string) {
	s.DB.Exec("UPDATE APP_USER SET PICTURE = ? WHERE ID = ?", pict, uid)
}
