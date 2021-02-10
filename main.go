package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func main() {
	// connect to database
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}