package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId    uint
	User      User
	ArticleId uint
	Article   Article
	Content   string
}
