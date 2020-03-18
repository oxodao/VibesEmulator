package models

import (
	"github.com/jinzhu/gorm"
)

/*
   private static final int FEMALE = 0;
   private static final int MALE = 1;

*/

// User Represent a user
type User struct {
	gorm.Model
	LastAction   int64  `json:"lastAction"`
	FirstName    string `json:"name"`
	Username     string `json:"username"`
	Gender       int    `json:"gender"`
	GenderWanted int    `json:"genderWanted"`
	Country      string `json:"country"`
	Age          int    `json:"age"`
	AgeFrom      int    `json:"ageFrom"`
	AgeTo        int    `json:"ageTo"`
	Picture      string `json:"picture"`
	Language     string `json:"gameLanguage"`
	IsPremium    bool   `json:"isPremium"`
	IsAdult      bool   `json:"isXRatedEnabled"`
	Password     string `json:"-"`
	LatestToken  string `json:"-"`
}

// GetUserWithPictureURL replaces the Picture field with a full link to the picture. Should not be saved in DB.
func (u *User) GetUserWithPictureURL() User {
	// SHOULD ABSOLUTELY NOT BE SAVED TO THE DB
	// If it is saved, that's maybe gorm's fault when updating the entity (If directly retreived from client)
	u.Picture = "https://vibes.oxodao.fr/pictures/" + u.Picture
	return *u
}
