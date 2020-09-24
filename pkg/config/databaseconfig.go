package config

import (
	"github.com/jinzhu/gorm"
	"log"
	"GoWebServiceWithCRUD/pkg/models"
)

var sqlConnectionString = "sqlserver://mgm-poc-user:super123@BCde@mgm-poc-db.database.windows.net:1433?database=mgm-poc-database&connection+timeout=30&charset=utf8mb4"
var dbType = "mssql"

func InitialMigration() {
	db, err := gorm.Open(dbType, sqlConnectionString)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	} else {
		log.Println("Connected to database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.User{})
}
