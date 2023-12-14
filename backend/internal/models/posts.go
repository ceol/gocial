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

func MockPost(id uint) Post {
	return Post{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    1,
	}
}
