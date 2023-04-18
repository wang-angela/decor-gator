package utils

import (
	"log"

	"github.com/decor-gator/backend/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var TX *gorm.DB
var err error

func InitDB(name string) {
	// Open data.db; if data does not exist, create it
	DB, err = gorm.Open(sqlite.Open("./pkg/databases/"+name+".db"), &gorm.Config{})
	// If there is an error, print error and message
	if err != nil {
		log.Print("Unable to connect to DB")
	}
	DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Image{})
}

func InitDBTest(name string) *gorm.DB {
	// Open data.db; if data does not exist, create it
	DB, err = gorm.Open(sqlite.Open("../databases/"+name+".db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	// If there is an error, print error and message
	if err != nil {
		log.Print("Unable to connect to DB")
	}

	TX = DB.Session(&gorm.Session{SkipDefaultTransaction: true})
	TX.AutoMigrate(&models.User{}, &models.Post{}, &models.Image{})

	return TX
}
