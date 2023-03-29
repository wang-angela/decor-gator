package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/decor-gator/backend/pkg/models"
	"github.com/decor-gator/backend/pkg/utils"
	"github.com/gorilla/mux"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var posts []models.Post

	// Prints an error if no posts are in the database.
	if utils.DB.Find(&posts).Error != nil {
		log.Printf("No posts exists.")
	}

	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var post models.Post

	// Prints error if the post doesn't exist.
	params := mux.Vars(r)["id"]
	utils.DB.First(&post, params)
	if post.ID == 0 {
		log.Fatalln("Post not found")
	}

	err = json.NewEncoder(w).Encode(&post)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Returns error if decoding is unsuccessful
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	if utils.DB.Create(&post).Error != nil {
		log.Println("Post already exists")
	}

	// Returns error if encoding is unsuccessful
	err = json.NewEncoder(w).Encode(&post)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	w.Header().Set("Content-Type", "application/json")

	// Search for user by id; id=0 if user not found
	params := mux.Vars(r)["id"]
	utils.DB.First(&post, params)
	if post.ID == 0 {
		log.Fatalln("Post not found")
	}

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	utils.DB.Save(&post)

	err = json.NewEncoder(w).Encode(&post)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var post models.Post
	params := mux.Vars(r)["id"]

	// Prints if deletion was not successful.
	if utils.DB.Delete(&post, params).Error != nil {
		log.Printf("Could not delete post.")
	}

	json.NewEncoder(w).Encode("You've successfully deleted this post.")
}
