package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	Email    string `gorm:"unique" json:"email"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var users []User

	// Prints an error if no users were in the data base.
	if db.Find(&users).Error != nil {
		log.Printf("There are no users in the database.")
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Printf("Error Encoding.")
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var user User

	// Prints an error id the user doesn't exists.
	params := mux.Vars(r)["id"]
	db.First(&user, params)
	if user.ID == 0 {
		log.Fatalln("User not found")
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
		log.Fatalln("User not found")
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalln("Error Decoding")
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

	// Prints if deletion was not successful.
	if db.Delete(&user, params).Error != nil {
		log.Printf("Could not delete user.")
	}

	json.NewEncoder(w).Encode("You've successfully deleted this user.")
}
