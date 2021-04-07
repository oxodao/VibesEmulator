package models

/*
   private static final int FEMALE = 0;
   private static final int MALE = 1;

*/

// User Represent a user
type User struct {
	ID           uint64 `json:"id" db:"ID"`
	CreatedAt    uint   `json:"-" db:"CREATED_AT"`
	LastAction   int64  `json:"lastAction" db:"LAST_ACTION"`
	FirstName    string `json:"name" db:"FIRSTNAME"`
	Username     string `json:"username" db:"USERNAME"`
	Gender       int    `json:"gender" db:"GENDER"`
	GenderWanted int    `json:"genderWanted" db:"GENDER_WANTED"`
	Country      string `json:"country" db:"COUNTRY"`
	Age          int    `json:"age" db:"AGE"`
	AgeFrom      int    `json:"ageFrom" db:"AGE_FROM"`
	AgeTo        int    `json:"ageTo" db:"AGE_TO"`
	Picture      string `json:"picture" db:"PICTURE"`
	Language     string `json:"gameLanguage" db:"LANG"`
	IsPremium    bool   `json:"isPremium"  db:"PREMIUM"`
	IsAdult      bool   `json:"isXRatedEnabled" db:"ADULT"`
	Password     string `json:"-" db:"PASSWORD"`
	LatestToken  string `json:"-" db:"LATEST_TOKEN"`
}

// GetUserWithPictureURL replaces the Picture field with a full link to the picture. Should not be saved in DB.
// This should be removed and translated as a DTO I think
func (u *User) GetUserWithPictureURL(webroot string) User {
	// SHOULD ABSOLUTELY NOT BE SAVED TO THE DB
	u.Picture = webroot + "/pictures/" + u.Picture
	return *u
}


func (u User) GetTableName() string {
	return "APP_USER"
}

func (u User) GetTableCreationScript() string {
	return `
		CREATE TABLE APP_USER (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			CREATED_AT INTEGER DEFAULT (cast(strftime('%s','now') as int)),
			LAST_ACTION INTEGER DEFAULT (cast(strftime('%s','now') as int)),
			FIRSTNAME VARCHAR(255),
			USERNAME VARCHAR(255),
			GENDER INTEGER,
			GENDER_WANTED INTEGER,
			COUNTRY VARCHAR(255),
			AGE INTEGER,
			AGE_FROM INTEGER,
			AGE_TO INTEGER,
			PICTURE VARCHAR(255),
			LANG VARCHAR(255),
			PREMIUM BOOL,
			ADULT BOOL,
			PASSWORD VARCHAR(255),
			LATEST_TOKEN VARCHAR(255)
		);
`
}