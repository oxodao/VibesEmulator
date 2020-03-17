package models

import "github.com/jinzhu/gorm"

// Picture represent an uploaded picture
type Picture struct {
	gorm.Model
	Name       string
	UploadedBy string
}
