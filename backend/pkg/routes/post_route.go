package routes

import (
	"github.com/decor-gator/backend/pkg/controllers"
	"github.com/gorilla/mux"
)

func PostRoutes(r *mux.Router) {
	r.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", controllers.GetPost).Methods("GET")
	r.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id}", controllers.UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", controllers.DeletePost).Methods("DELETE")
}
