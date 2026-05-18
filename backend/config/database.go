package config

import (
	"backend/models"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Can not connect to the database:", err)
	}

	database.AutoMigrate(&models.Book{})
	DB = database
	log.Println("Database connection successful")
}
