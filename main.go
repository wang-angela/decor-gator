package main

import (
	_ "encoding/json"
	_ "log"
	_ "net/http"

	_ "github.com/gorilla/mux"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
