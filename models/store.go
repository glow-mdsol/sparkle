package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Database(dbType, dbConnection string) *gorm.DB {
	db, err := gorm.Open(dbType, dbConnection)
	if err != nil {
		fmt.Printf("Error connecting to DB: <%s> \n", err)

		panic(fmt.Errorf("failed to connect database with error  <%s> \n", err))
	}
	defer func(){
		err := db.Close()
		if err != nil {

		}
	}()

	// Migrate the properties
	db.AutoMigrate(&Namespace{})
	db.AutoMigrate(&Property{})

	return db
}

