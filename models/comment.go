package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId    uint64
	User      User
	ArticleId uint64
	Article   Article
	Content   string
}
