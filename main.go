package main

import (
	"log"

	"github.com/herusdianto/gorm_crud_example/database"
	"github.com/herusdianto/gorm_crud_example/models"
)

func main() {
	// database configs
	dbUser, dbPassword, dbName := "root", "root", "gorm_crud_example"

	db, err := database.ConnectToDB(dbUser, dbPassword, dbName)

	// unable to connect to database
	if err != nil {
		log.Fatalln(err)
	}

	// ping to database
	err = db.DB().Ping()

	// error ping to database
	if err != nil {
		log.Fatalln(err)
	}

	// migration
	db.AutoMigrate(&models.Contact{})

	defer db.Close()
}
