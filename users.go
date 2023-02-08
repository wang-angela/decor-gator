package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Email    string `gorm:"unique" json:"email"`
}

func initMigration() {
	// Open data.db; if data does not exist, create it
	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	// If there is an error, print error and message
	if err != nil {
		log.Print("Unable to connect to DB")
	}
	db.AutoMigrate(&User{})
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var user User

	params := mux.Vars(r)["id"]
	db.First(&user, params)
	if user.ID == 0 {
		log.Println("User not found")
	}

	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Returns error if decoding is unsuccessful
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	user.Password = encrypt(user.Password)
	if db.Create(&user).Error != nil {
		log.Println("User already exists")
	}

	// Returns error if encoding is unsuccessful
	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	w.Header().Set("Content-Type", "application/json")

	// Search for user by id; id=0 if user not found
	params := mux.Vars(r)["id"]
	db.First(&user, params)
	if user.ID == 0 {
		log.Println("User not found")
	}

	// Retrieve and store current
	cur := user.Password

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	// If password didn't change, keep current hash
	if comparePassword(user.Password, cur) {
		user.Password = cur
	} else {
		user.Password = encrypt(user.Password)
	}

	db.Save(&user)

	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var user User
	params := mux.Vars(r)["id"]
	db.Delete(&user, params)
	json.NewEncoder(w).Encode("You've successfully deleted this user.")
}
