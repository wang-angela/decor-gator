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

var tx *gorm.DB

func initDB() {
	// Open data.db; if data does not exist, create it
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	// If there is an error, print error and message
	if err != nil {
		log.Print("Unable to connect to DB")
	}

	tx = db.Session(&gorm.Session{SkipDefaultTransaction: true})
	tx.AutoMigrate(&User{}, &Post{})
}

func deleteLast() {
	db.Exec("DELETE FROM users WHERE id = (SELECT MAX(id) FROM users)")
	db.Exec("DELETE FROM posts")
}

func TestGetAllUsers(t *testing.T) {
	initDB()

	// Send new request with json body info
	req, err := http.NewRequest("POST", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUsers)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp []User
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Errorf("Invalid response, expected list of users, got %v", rr.Body.String())
	}

	if len(resp) < 1 {
		t.Errorf("Invalid number of users, expected 1, got %v", len(resp))
	}
}

func TestGetUser(t *testing.T) {
	initDB()

	// Send new request with json body info
	req, err := http.NewRequest("GET", "/users/simon@simonkurt.com", nil)
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

	if resp["username"] != "simon.kurt" {
		t.Errorf("Username is invalid, expected simon.kurt, got %v", resp["username"])
	}
}

func TestCreateUser(t *testing.T) {
	initDB()
	tx.SavePoint("sp1")

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

	tx.RollbackTo("sp1")
}

func TestUpdateUser(t *testing.T) {
	initDB()

	// Request Body
	jsonBody := []byte(`{
		"username": "billy.scott",
		"password": "123abc",
		"email":    "will.scott@thehouse.com"
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	// Send new request with json body info
	req, err := http.NewRequest("PUT", "/users/will.scott@thehouse.com", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/users/{email}", updateUser)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp["username"] != "billy.scott" {
		t.Errorf("Username is invalid, expected billy.scott, got %v", resp["username"])
	}
}

func TestDeleteUser(t *testing.T) {
	initDB()

	// Send new request with json body info
	req, err := http.NewRequest("DELETE", "/users/will.scott@thehouse.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/users/{email}", deleteUser)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp["deleted_at"] == "" {
		t.Errorf("Has not been deleted")
	}
}
