package fixtures

import (
	"fmt"
	"github.com/oxodao/vibes/models"
	"github.com/oxodao/vibes/services"
)

func InitializeDb(prv *services.Provider) {
	tables := []models.DbModel{
		models.Picture{},
		models.User{},
		models.Contact{},
		models.Message{},
		models.Question{},
		models.Answer{},
	}

	fmt.Println("Dropping old tables...")
	for i := len(tables) - 1; i >= 0; i-- {
		_, e := prv.DB.Exec(`DROP TABLE IF EXISTS ` + tables[i].GetTableName())
		if e != nil {
			panic(e)
		}
	}

	for _, t := range tables {
		_, e := prv.DB.Exec(t.GetTableCreationScript())
		if e != nil {
			fmt.Println(t.GetTableName())
			panic(e)
		}
		fmt.Printf("\t- Table %v created.\n", t.GetTableName())
	}

}
