package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	initMigration()
	initRouter()
}

func initMigration() {
	// Open data.db; if data does not exist, create it
	db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	// If there is an error, print error and message
	if err != nil {
		log.Print("Unable to connect to DB")
	}
	db.AutoMigrate(&User{}, &Post{})
}

func initRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	r.HandleFunc("/posts", getPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", getPost).Methods("GET")
	r.HandleFunc("/posts", createPost).Methods("POST")
	r.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}