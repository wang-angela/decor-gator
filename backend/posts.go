package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey" json:"id"`
	Title         string `json:"title"`
	FurnitureType string `json:"furnitureType"`
	UserPosted    string `json:"userPosted"` // Change later to user object once we learn how to test that on Postman.
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []Post

	// Prints an error if no posts are in the database.
	if db.Find(&posts).Error != nil {
		log.Printf("No posts exists.")
	}

	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post

	// Prints error if the post doesn't exist.
	params := mux.Vars(r)["id"]
	db.First(&post, params)
	if post.ID == 0 {
		log.Fatalln("Post not found")
	}

	err = json.NewEncoder(w).Encode(&post)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Returns error if decoding is unsuccessful
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	if db.Create(&post).Error != nil {
		log.Println("Post already exists")
	}

	// Returns error if encoding is unsuccessful
	err = json.NewEncoder(w).Encode(&post)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	w.Header().Set("Content-Type", "application/json")

	// Search for user by id; id=0 if user not found
	params := mux.Vars(r)["id"]
	db.First(&post, params)
	if post.ID == 0 {
		log.Fatalln("Post not found")
	}

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	db.Save(&post)

	err = json.NewEncoder(w).Encode(&post)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	params := mux.Vars(r)["id"]

	// Prints if deletion was not successful.
	if db.Delete(&post, params).Error != nil {
		log.Printf("Could not delete post.")
	}

	json.NewEncoder(w).Encode("You've successfully deleted this post.")
}
