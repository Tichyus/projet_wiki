package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username	string
	Password	uint
	Articles	[]Article
	Comments 	[]Comment
}

type Article struct {
	gorm.Model
	UserId		uint
	User			User
	Title			string
	Content		string
	Comments 	[]Comment
}

type Comment struct {
	gorm.Model
	UserId		uint
	User			User
	ArticleId	uint
	Article		Article
	Content		string
}