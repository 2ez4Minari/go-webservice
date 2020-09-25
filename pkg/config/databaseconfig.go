package config

import (
	"GoWebServiceWithCRUD/pkg/models"
	"github.com/jinzhu/gorm"
	"log"
)

var SqlConnectionString = "sqlserver://mgm-poc-user:super123@BCde@mgm-poc-db.database.windows.net:1433?database=mgm-poc-database&connection+timeout=30&charset=utf8mb4"
var DbType = "mssql"

func InitialMigration() {
	db, err := gorm.Open(DbType, SqlConnectionString)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	} else {
		log.Println("Connected to database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Customer{})
}
