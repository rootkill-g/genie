package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"rootkill/genie/models"
)

var DB *gorm.DB

// DB Migration Schema

func InitDb() (*gorm.DB, error) {
	// Initialize Database
	fmt.Println("Connecting to database")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	// Migrate the schema
	fmt.Println("Running database migrations")
	db.AutoMigrate(&models.User{})

	DB = db
	return DB, nil
}
