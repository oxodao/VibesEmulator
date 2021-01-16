package dal

import "github.com/oxodao/vibes/services"

/**
@TODO Handle errors
**/

// UpdateAge blabla
func UpdateAge(prv *services.Provider, uid uint64, age int, ageFrom int, ageTo int) {
	prv.DB.Exec(`
		UPDATE APP_USER
		SET
			AGE = CASE WHEN $1 <> -1 THEN $1 ELSE AGE END,
			AGE_FROM = CASE WHEN $2 <> -1 THEN $2 ELSE AGE END,
			AGE_TO = CASE WHEN $3 <> -1 THEN $3 ELSE AGE END
		WHERE ID = $4
	`, age, ageFrom, ageTo, uid)
}

// UpdateFirstname blabla
func UpdateFirstname(prv *services.Provider, uid uint64, name string) {
	prv.DB.Exec("UPDATE APP_USER SET FIRSTNAME = ? WHERE ID = ?", name, uid)
}

// UpdateGender blabla
func UpdateGender(prv *services.Provider, uid uint64, gender int, genderWanted int) {
	prv.DB.Exec(`
		UPDATE APP_USER
		SET
			GENDER = CASE WHEN $1 <> -1 THEN $1 ELSE GENDER END,
			GENDER_WANTED = CASE WHEN $2 <> -1 THEN $2 ELSE GENDER_WANTED END
		WHERE ID = $3
	`, gender, genderWanted, uid)
}

// UpdateAdult blabla
func UpdateAdult(prv *services.Provider, uid uint64, adult bool) {
	prv.DB.Exec("UPDATE APP_USER SET ADULT = ? WHERE ID = ?", adult, uid)
}

// UpdateLanguage blabla
func UpdateLanguage(prv *services.Provider, uid uint64, lang string) {
	prv.DB.Exec("UPDATE APP_USER SET LANG = ? WHERE ID = ?", lang, uid)
}

// UpdatePicture blabla
func UpdatePicture(prv *services.Provider, uid uint64, pict string) {
	prv.DB.Exec("UPDATE APP_USER SET PICTURE = ? WHERE ID = ?", pict, uid)
}
