package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password uint64
	Articles []Article
	Comments []Comment
}
