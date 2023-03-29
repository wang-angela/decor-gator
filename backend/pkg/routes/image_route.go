package routes

import (
	"github.com/decor-gator/backend/pkg/controllers"
	"github.com/gorilla/mux"
)

func ImageRoutes(r *mux.Router) {
	r.HandleFunc("/images", controllers.GetImages).Methods("GET")
	r.HandleFunc("/images/{id}", controllers.GetImage).Methods("GET")
	r.HandleFunc("/images", controllers.CreateImage).Methods("POST")
	r.HandleFunc("/images/{id}", controllers.UpdateImage).Methods("PUT")
	r.HandleFunc("/images/{id}", controllers.DeleteImage).Methods("DELETE")
}
