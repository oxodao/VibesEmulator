package models

import "time"

/*
   private static final int FEMALE = 0;
   private static final int MALE = 1;

*/

// User Represent a user
type User struct {
	ID           uint      `json:"-" db:"ID"`
	CreatedAt    time.Time `json:"-" db:"CREATED_AT"`
	LastAction   int64     `json:"lastAction" db:"LAST_ACTION"`
	FirstName    string    `json:"name" db:"FIRSTNAME"`
	Username     string    `json:"username" db:"USERNAME"`
	Gender       int       `json:"gender" db:"GENDER"`
	GenderWanted int       `json:"genderWanted" db:"GENDER_WANTED"`
	Country      string    `json:"country" db:"COUNTRY"`
	Age          int       `json:"age" db:"AGE"`
	AgeFrom      int       `json:"ageFrom" db:"AGE_FROM"`
	AgeTo        int       `json:"ageTo" db:"AGE_TO"`
	Picture      string    `json:"picture" db:"PICTURE"`
	Language     string    `json:"gameLanguage" db:"LANG"`
	IsPremium    bool      `json:"isPremium"  db:"PREMIUM"`
	IsAdult      bool      `json:"isXRatedEnabled" db:"ADULT"`
	Password     string    `json:"-" db:"PASSWORD"`
	LatestToken  string    `json:"-" db:"LATEST_TOKEN"`
}

// GetUserWithPictureURL replaces the Picture field with a full link to the picture. Should not be saved in DB.
func (u *User) GetUserWithPictureURL() User {
	// SHOULD ABSOLUTELY NOT BE SAVED TO THE DB
	// If it is saved, that's maybe gorm's fault when updating the entity (If directly retreived from client)
	u.Picture = "https://vibes.oxodao.fr/pictures/" + u.Picture
	return *u
}
