package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetAllPosts(t *testing.T) {
	initDB()

	// Send new request with json body info
	req, err := http.NewRequest("POST", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getPosts)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp []Post
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Errorf("Invalid response, expected list of posts, got %v", rr.Body.String())
	}

	if len(resp) < 1 {
		t.Errorf("Invalid number of posts, expected 1, got %v", len(resp))
	}
}

func TestGetPost(t *testing.T) {
	initDB()

	// Send new request with json body info
	req, err := http.NewRequest("GET", "/posts/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/posts/{id}", getPost)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp["title"] != "Sofa for sale!" {
		t.Errorf("Posts is invalid, expected \"Sofa for sale\", got %v", resp["title"])
	}
}

func TestCreatePost(t *testing.T) {
	initDB()
	tx.SavePoint("sp1")

	// Request Body
	jsonBody := []byte(`{
		"id": 4,
		"title": "Hello Kitty Chair for Rent!",
		"furnitureType": "Chair",
		"userPosted": "john.smith"
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	// Send new request with json body info
	req, err := http.NewRequest("POST", "/posts", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createPost)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp["title"] != "Hello Kitty Chair for Rent!" {
		t.Errorf("Post title is invalid, expected \"Hello Kitty Chair for Rent!\", got %v", resp["title"])
	}

	tx.RollbackTo("sp1")
}

func TestUpdatePost(t *testing.T) {
	initDB()
	tx.SavePoint("sp2")

	// Request Body
	jsonBody := []byte(`{
		"id": 2,
		"title": "Selling a Dining Table",
		"furnitureType": "Table",
		"userPosted": "william-scott"
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	// Send new request with json body info
	req, err := http.NewRequest("PUT", "/posts/2", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/posts/{id}", updatePost)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp["title"] != "Selling a Dining Table" {
		t.Errorf("Post Title is invalid, expected \"Selling a Dining Table\", got %v", resp["title"])
	}

	tx.RollbackTo("sp2")
}

func TestDeletePost(t *testing.T) {
	initDB()
	tx.SavePoint("sp3")

	// Send new request with json body info
	req, err := http.NewRequest("DELETE", "/posts/3", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/posts/{id}", deletePost)
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

	tx.RollbackTo("sp3")
}
