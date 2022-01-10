package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectToDB(dbUser string, dbPassword string, dbName string) (*gorm.DB, error) {
	var connectionString = fmt.Sprintf(
		"host='localhost' port=5432 user='%s' password='%s' dbname='%s' sslmode=disable",
		dbUser, dbPassword, dbName,
	)

	return gorm.Open("postgres", connectionString)
}
