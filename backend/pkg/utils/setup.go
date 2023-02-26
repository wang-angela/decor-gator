package utils

import (
	"log"

	"github.com/decor-gator/backend/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDB() {
	// Open data.db; if data does not exist, create it
	DB, err = gorm.Open(sqlite.Open("./pkg/models/data.db"), &gorm.Config{})
	// If there is an error, print error and message
	if err != nil {
		log.Print("Unable to connect to DB")
	}
	DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Image{})
}
