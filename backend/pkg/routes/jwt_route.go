package routes

import (
	"github.com/decor-gator/backend/pkg/controllers"
	"github.com/gorilla/mux"
)

func JwtRoutes(r *mux.Router) {
	r.HandleFunc("/authenticate", controllers.CreateTokenEndpoint).Methods("POST")
	r.HandleFunc("/protected", controllers.ProtectedEndpoint).Methods("GET")
	r.HandleFunc("/test", controllers.ValidateMiddleware(controllers.TestEndpoint)).Methods("GET")
}
