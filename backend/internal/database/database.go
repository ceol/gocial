package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	DB = db
	return err
}
