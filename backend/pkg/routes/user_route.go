package routes

import (
	"github.com/decor-gator/backend/pkg/controllers"
	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc("/user", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/user/{email}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/user/{email}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{email}", controllers.DeleteUser).Methods("DELETE")
}
