package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	imageID       int    `gorm:"primaryKey" json:"id"`
	imageByteData string `json:"ByteData"`
}

func getImages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var images []Image

	// Prints an error if no posts are in the database.
	if db.Find(&images).Error != nil {
		log.Printf("No images exists.")
	}

	err = json.NewEncoder(w).Encode(images)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func getImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var image Image

	// Prints error if the post doesn't exist.
	params := mux.Vars(r)["id"]
	db.First(&image, params)
	if image.ID == 0 {
		log.Fatalln("Image not found")
	}

	err = json.NewEncoder(w).Encode(&image)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func createImage(w http.ResponseWriter, r *http.Request) {
	var image Image
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Returns error if decoding is unsuccessful
	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	if db.Create(&image).Error != nil {
		log.Println("Image already exists")
	}

	// Returns error if encoding is unsuccessful
	err = json.NewEncoder(w).Encode(&image)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func updateImage(w http.ResponseWriter, r *http.Request) {
	var image Image
	w.Header().Set("Content-Type", "application/json")

	// Search for user by id; id=0 if user not found
	params := mux.Vars(r)["id"]
	db.First(&image, params)
	if image.ID == 0 {
		log.Fatalln("Image not found")
	}

	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	db.Save(&image)

	err = json.NewEncoder(w).Encode(&image)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func deleteImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var image Image
	params := mux.Vars(r)["id"]

	// Prints if deletion was not successful.
	if db.Delete(&image, params).Error != nil {
		log.Printf("Could not delete image.")
	}

	json.NewEncoder(w).Encode("You've successfully deleted this image.")
}
