package database

import (
	"server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("../gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.ExchangeRate{})
	return db, nil
}
