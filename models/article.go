package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	UserId   uint
	User     User
	Title    string
	Content  string
	Comments []Comment
}
