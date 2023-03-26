package routes

import (
	"github.com/decor-gator/backend/pkg/controllers"
	"github.com/gorilla/mux"
)

func EmailRoutes(r *mux.Router) {
	r.HandleFunc("/emails/ForgotPassword", controllers.UpdateImage).Methods("PUT")
}
