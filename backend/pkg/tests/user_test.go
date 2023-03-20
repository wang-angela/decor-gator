package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/decor-gator/backend/pkg/controllers"
	"github.com/decor-gator/backend/pkg/models"
	"github.com/decor-gator/backend/pkg/utils"
	"github.com/gorilla/mux"
)

func TestGetAllUsers(t *testing.T) {
	utils.InitDBTest("test")

	// Send new request with json body info
	req, err := http.NewRequest("POST", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetUsers)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp []models.User
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Errorf("Invalid response, expected list of users, got %v", rr.Body.String())
	}

	if len(resp) < 1 {
		t.Errorf("Invalid number of users, expected 1, got %v", len(resp))
	}
}

func TestGetUser(t *testing.T) {
	utils.InitDBTest("test")

	// Send new request with json body info
	req, err := http.NewRequest("GET", "/users/simon@simonkurt.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/users/{email}", controllers.GetUser)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp["username"] != "simon-kurt" {
		t.Errorf("Username is invalid, expected simon-kurt, got %v", resp["username"])
	}
}

func TestCreateUser(t *testing.T) {
	TX := utils.InitDBTest("test")
	TX.SavePoint("sp1")

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
	handler := http.HandlerFunc(controllers.CreateUser)

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

	TX.RollbackTo("sp1")
}

func TestUpdateUser(t *testing.T) {
	TX := utils.InitDBTest("test")
	TX.SavePoint("sp2")

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
	r.HandleFunc("/users/{email}", controllers.UpdateUser)
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

	TX.RollbackTo("sp2")
}

func TestDeleteUser(t *testing.T) {
	TX := utils.InitDBTest("test")
	TX.SavePoint("sp3")

	// Send new request with json body info
	req, err := http.NewRequest("DELETE", "/users/will.scott@thehouse.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/users/{email}", controllers.DeleteUser)
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

	TX.RollbackTo("sp3")
}
