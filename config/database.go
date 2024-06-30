package config

import (
	"log"

	"github.com/glebarez/sqlite"
	"github.com/thrtn85/library-mgmt/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Connect to SQLite database
	database, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	
	// Automatically merge schema
	database.AutoMigrate(&models.Book{}, &models.User{})
	
	// Assign the database connection to the global DB variable
	DB = database
}