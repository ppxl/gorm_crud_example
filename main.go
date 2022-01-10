package main

import (
	"log"

	"github.com/ppxl/gorm_crud_example/configs"
	"github.com/ppxl/gorm_crud_example/database"
	"github.com/ppxl/gorm_crud_example/models"
	"github.com/ppxl/gorm_crud_example/repositories"
)

func main() {
	// database configs
	dbUser, dbPassword, dbName := "postgres", "postgres", "mydb"

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

	contactRepository := repositories.NewContactRepository(db)

	route := configs.SetupRoutes(contactRepository)

	route.Run(":8000")
}
