package main

import (
	_ "encoding/json"
	_ "log"
	_ "net/http"

	_ "github.com/gorilla/mux"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
