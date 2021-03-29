package database

import (
	"projet_wiki/models"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DbConn *gorm.DB
)

func Connect() error {
	dsn := "root:root@tcp(127.0.0.1:3306)/flamingo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err != nil {
		panic("failed to connect database")
	}

	DbConn = db

	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	return nil
}
