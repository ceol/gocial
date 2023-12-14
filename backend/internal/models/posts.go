package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    uint
	User      User
	Content   string
}
