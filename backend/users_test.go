package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() {
	// Open data.db; if data does not exist, create it
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// If there is an error, print error and message
	if err != nil {
		log.Print("Unable to connect to DB")
	}
	db.AutoMigrate(&User{}, &Post{})
}

func clearDB() {
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM posts")
}

func TestCreateUser(t *testing.T) {
	initDB()
	clearDB()

	// Request Body
	jsonBody := []byte(`{
		"username": "john.smith",
		"password": "123abc",
		"email":    "john.smith@gmail.com"
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	// Send new request with json body info
	req, err := http.NewRequest("POST", "/users", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp["username"] != "john.smith" {
		t.Errorf("username is invalid, expected john.smith, got %v", resp["username"])
	}
}

func TestGetUser(t *testing.T) {
	initDB()

	// Send new request with json body info
	req, err := http.NewRequest("GET", "/users/john.smith@gmail.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/users/{email}", getUser)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp["username"] != "john.smith" {
		t.Errorf("username is invalid, expected john.smith, got %v", resp["username"])
	}
}
