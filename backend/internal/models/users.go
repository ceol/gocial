package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	UserName     string
	Email        string
	PasswordHash string
}
