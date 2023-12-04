package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID  uint
	User    User
	Content string
}