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
	dsn := "root:@tcp(127.0.0.1:3306)/flamingo?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err != nil {
		panic("failed to connect database")
	}

	DbConn = db

	// Creates the tables, missing foreign keys, constraints, columns and indexes for the specified models
	db.AutoMigrate(&models.User{}, &models.Article{}, &models.Comment{})

	return nil
}
