package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/decor-gator/backend/pkg/configs"
	"github.com/decor-gator/backend/pkg/controllers"
	"github.com/decor-gator/backend/pkg/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestGetAllUsers(t *testing.T) {
	configs.ConnectDB()

	// Send new request with json body info
	req, err := http.NewRequest("GET", "/user", nil)
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
	configs.ConnectDB()

	// Send new request with json body info
	req, err := http.NewRequest("GET", "/user/john.smith@hotmail.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/user/{email}", controllers.GetUser)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp["username"] != "BasicGuy" {
		t.Errorf("Username is invalid, expected BasicGuy, got %v", resp["username"])
	}
}

func TestCreateUser(t *testing.T) {
	configs.ConnectDB()

	// Request Body
	jsonBody := []byte(`{
		"firstName": "Angela",
		"lastName": "Wang",
		"email":    "john.smith@gmail.com",
		"username": "iLuvGophers",
		"password": "golang!!!"
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	// Send new request with json body info
	req, err := http.NewRequest("POST", "/user", bodyReader)
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

	fmt.Println(resp)

	if resp["InsertID"] != nil {
		t.Errorf("Insert failed")
	}

	configs.GetCollection(configs.DB, "users").DeleteOne(context.TODO(),
		bson.D{{Key: "username", Value: "iLuvGophers"}},
	)
}

func TestUpdateUser(t *testing.T) {
	// Request Body
	jsonBody := []byte(`{
		"username": "BasicGal",
		"password": "myNewPassword!"
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	// Send new request with json body info
	req, err := http.NewRequest("PUT", "/user/jane.doe@gmail.com", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/user/{email}", controllers.UpdateUser)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp *mongo.UpdateResult
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp.MatchedCount != 1 {
		t.Errorf("Fail to update")
	}

	// Reversing changes
	update := bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "first_name", Value: "Jane"},
			{Key: "last_name", Value: "Doe"},
			{Key: "email", Value: "jane.doe@gmail.com"},
			{Key: "username", Value: "BasicGal"},
			{Key: "password", Value: "little-lamb"},
		},
	}}

	configs.GetCollection(configs.DB, "users").UpdateOne(context.TODO(),
		bson.D{{Key: "email", Value: "jane.doe@gmail.com"}},
		update,
	)
}

func TestDeleteUser(t *testing.T) {
	user := &models.User{
		ID:        primitive.NewObjectID(),
		FirstName: "Jack",
		LastName:  "Harrison",
		Email:     "jackieboi@mail.com",
		Username:  "jacksonSon",
		Password:  "SuperPass",
	}

	configs.GetCollection(configs.DB, "users").InsertOne(context.TODO(), user)

	// Send new request with json body info
	req, err := http.NewRequest("DELETE", "/user/jackieboi@mail.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/user/{email}", controllers.DeleteUser)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp *mongo.DeleteResult
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp.DeletedCount != 1 {
		t.Errorf("Has not been deleted")
	}
}
