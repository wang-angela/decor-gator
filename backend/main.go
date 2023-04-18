package main

import (
	"log"
	"net/http"

	"github.com/decor-gator/backend/pkg/controllers"
	"github.com/decor-gator/backend/pkg/routes"
	"github.com/decor-gator/backend/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	utils.InitDB("data")
	controllers.Init()

	// Routes
	routes.UserRoutes(r)
	routes.PostRoutes(r)
	routes.ImageRoutes(r)
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
