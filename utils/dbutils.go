package utils

import (
	"github.com/jinzhu/gorm"
	"github.com/oxodao/vibes/services"
)

// GetRandomFromDB returns the expression for random in the current DB
func GetRandomFromDB(prv *services.Provider) *gorm.SqlExpr {
	s := prv.DB.Dialect().GetName()

	if s == "mysql" {
		return gorm.Expr("rand()")
	}

	if s == "sqlite3" || s == "postgresql" {
		return gorm.Expr("random()")
	}

	return nil
}