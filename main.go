package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initMigration()
	initRouter()
}

func initRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}
