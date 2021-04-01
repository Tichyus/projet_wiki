package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	UserID   uint64
	User     User
	Title    string
	Content  string
	Comments []Comment
}
