package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	db, err := gorm.Open(sqlite.Open("gocial.db"), &gorm.Config{})
	DB = db
	return err
}
