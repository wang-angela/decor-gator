package main

import (
	"log"
	"net/http"

	"github.com/decor-gator/backend/pkg/routes"
	"github.com/decor-gator/backend/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	utils.InitDB()

	routes.UserRoutes(r)
	routes.PostRoutes(r)
	routes.ImageRoutes(r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
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
