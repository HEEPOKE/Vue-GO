package database

import (
	"Server/Api/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	dsn := os.Getenv("DB_USER") + ":@tcp(" + os.Getenv("HOST") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
}
