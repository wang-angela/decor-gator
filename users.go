package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func initMigration() {
	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	// If there is an error, print error and message
	if err != nil {
		log.Print("Unable to connect to DB")
	}
	db.AutoMigrate(&User{})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	db.Create(&user)

	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}
