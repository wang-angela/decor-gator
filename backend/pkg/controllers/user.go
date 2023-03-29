/* Code referenced from:
* https://www.youtube.com/watch?v=KPftgI40WHI by the Daily Code Buffer
* https://blog.logrocket.com/routing-go-gorilla-mux/ by Paul Akinyemi
 */

package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/decor-gator/backend/pkg/models"
	"github.com/decor-gator/backend/pkg/utils"
	"github.com/gorilla/mux"
)

var err error

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var users []models.User

	// Prints an error if no users were in the data base.
	if utils.DB.Find(&users).Error != nil {
		log.Printf("There are no users in the database.")
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Printf("Error Encoding.")
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var user models.User

	// Prints an error id the user doesn't exists.
	params := mux.Vars(r)["email"]
	utils.DB.Where("email = ?", params).First(&user)
	if user.Email == "" {
		log.Println("User not found")
	}

	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Returns error if decoding is unsuccessful
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	user.Password = utils.Encrypt(user.Password)
	if utils.DB.Create(&user).Error != nil {
		log.Println("User already exists")
		return
	}

	email := []string{user.Email}
	SendWelcomeEmail(email)

	// Returns error if encoding is unsuccessful
	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")

	// Search for user by id; id=0 if user not found
	params := mux.Vars(r)["email"]
	utils.DB.Where("email = ?", params).First(&user)
	if user.Email == "" {
		log.Println("User not found")
	}

	// Retrieve and store current
	cur := user.Password

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	// If password didn't change, keep current hash
	if utils.ComparePassword(user.Password, cur) {
		user.Password = cur
	} else {
		user.Password = utils.Encrypt(user.Password)
	}

	utils.DB.Save(&user)

	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var user models.User
	params := mux.Vars(r)["email"]

	utils.DB.Where("email = ?", params).Delete(&user)

	// Prints if deletion was not successful.
	if user.Email != "" {
		log.Printf("Could not delete user.")
	}

	json.NewEncoder(w).Encode("You've successfully deleted this user.")
}
