package routes

import (
	"github.com/decor-gator/backend/pkg/controllers"
	"github.com/gorilla/mux"
)

func BucketRoutes(r *mux.Router) {
	r.HandleFunc("/buckets/{filename}", controllers.UploadObjectHelper).Methods("POST")
	r.HandleFunc("/buckets", controllers.ListBucketsHelper).Methods("GET")
	r.HandleFunc("/buckets/{filename}", controllers.DeleteObjectHelper).Methods("DELETE")
}
