package main

import (
	"log"
	"net/http"

	"github.com/decor-gator/backend/pkg/routes"
	"github.com/decor-gator/backend/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	SendWelcomeEmail([]string{"sgallic5@gmail.com"})
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
	db.AutoMigrate(&User{}, &Post{}, &Image{})
}

func initRouter() {
	r := mux.NewRouter()
	utils.InitDB("data")

	routes.UserRoutes(r)
	routes.PostRoutes(r)
	routes.ImageRoutes(r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

/*
func InitRouter() {

	r.HandleFunc("/login", Login)
	r.HandleFunc("/home", Home)
	r.HandleFunc("/refresh", Refresh)
}
*/
