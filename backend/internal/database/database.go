package database

import (
	"github.com/ceol/gocial/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Open(uri string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(uri), &gorm.Config{})
}

func Connect(uri string) error {
	db, err := Open(uri)
	DB = db
	return err
}

func Disconnect() error {
	db, _ := DB.DB()
	return db.Close()
}

func Migrate() []error {
	if DB == nil {
		panic("No database connection")
	}

	var errors []error

	for _, model_ptr := range models.Models {
		errors = append(errors, DB.AutoMigrate(model_ptr))
	}

	return errors
}

func DropTables() []error {
	if DB == nil {
		panic("No database connection")
	}

	var errors []error

	for _, model_ptr := range models.Models {
		errors = append(errors, DB.Migrator().DropTable(model_ptr))
	}

	return errors
}
