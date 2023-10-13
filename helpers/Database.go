package helpers

import (
	"kp-elibrary-golang/config"
	"kp-elibrary-golang/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() {
	dsn := config.DbConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&models.Role{}, &models.User{}, &models.BookCategory{}, &models.Book{}, &models.Category{})
	DB = db
}
