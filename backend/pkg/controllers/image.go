package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/decor-gator/backend/pkg/models"
	"github.com/decor-gator/backend/pkg/utils"
	"github.com/gorilla/mux"
)

func GetImages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var images []models.Image

	// Prints an error if no posts are in the database.
	if utils.DB.Find(&images).Error != nil {
		log.Printf("No images exists.")
	}

	err = json.NewEncoder(w).Encode(images)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var image models.Image

	// Prints error if the post doesn't exist.
	params := mux.Vars(r)["id"]
	utils.DB.First(&image, params)
	if image.ID == 0 {
		log.Fatalln("Image not found")
	}

	err = json.NewEncoder(w).Encode(&image)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func CreateImage(w http.ResponseWriter, r *http.Request) {
	var image models.Image
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Returns error if decoding is unsuccessful
	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	if utils.DB.Create(&image).Error != nil {
		log.Println("Image already exists")
	}

	// Returns error if encoding is unsuccessful
	err = json.NewEncoder(w).Encode(&image)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func UpdateImage(w http.ResponseWriter, r *http.Request) {
	var image models.Image
	w.Header().Set("Content-Type", "application/json")

	// Search for user by id; id=0 if user not found
	params := mux.Vars(r)["id"]
	utils.DB.First(&image, params)
	if image.ID == 0 {
		log.Fatalln("Image not found")
	}

	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	utils.DB.Save(&image)

	err = json.NewEncoder(w).Encode(&image)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func DeleteImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var image models.Image
	params := mux.Vars(r)["id"]

	// Prints if deletion was not successful.
	if utils.DB.Delete(&image, params).Error != nil {
		log.Printf("Could not delete image.")
	}

	json.NewEncoder(w).Encode("You've successfully deleted this image.")
}
