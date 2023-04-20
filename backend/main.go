package main

import (
	"log"
	"net/http"

	"github.com/decor-gator/backend/pkg/configs"
	"github.com/decor-gator/backend/pkg/controllers"
	"github.com/decor-gator/backend/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env")
	}

	r := mux.NewRouter()

	// Connect database
	configs.ConnectDB()
	controllers.InitAWSSession()
	controllers.CreateBucket()

	// Routes
	routes.UserRoutes(r)
	routes.PostRoutes(r)
	routes.EmailRoutes(r)
	routes.JwtRoutes(r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
