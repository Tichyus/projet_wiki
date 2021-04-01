package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID    uint64
	User      User
	ArticleID uint64
	Article   Article
	Content   string
}
