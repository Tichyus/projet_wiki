package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
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
	User		User
	Title		string
	Content		string
	Comments 	[]Comment
  }

  type Comment struct {
	gorm.Model
	UserId		uint
	User		User
	ArticleId	uint
	Article		Article
	Content		string
  }

func main() {
	// connect database
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}