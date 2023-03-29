package routes

import (
	"github.com/decor-gator/backend/pkg/controllers"
	"github.com/gorilla/mux"
)

func EmailRoutes(r *mux.Router) {
	r.HandleFunc("/emails/{email}", controllers.HelperForgotPassword).Methods("PUT")
}
