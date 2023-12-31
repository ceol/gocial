package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	UserName     string
	Email        string
	PasswordHash string
}
