package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetAllImages(t *testing.T) {
	initDB()

	// Send new request with json body info
	req, err := http.NewRequest("GET", "/images", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getImages)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp []Image
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Errorf("Invalid response, expected list of images, got %v", rr.Body.String())
	}

	if len(resp) < 1 {
		t.Errorf("Invalid number of images, expected 1, got %v", len(resp))
	}
}

func TestGetImage(t *testing.T) {
	initDB()

	// Send new request with json body info
	req, err := http.NewRequest("GET", "/images/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/images/{id}", getImage)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &resp)

	const mockData = "joewjoiu432987(*&(*$#&*(Wuroiesjfkjaskjdoiwj"
	if resp["ByteData"] != mockData {
		t.Errorf("Data is invalid, expected \"joewjoiu432987(*&(*$#&*(Wuroiesjfkjaskjdoiwj\" got %v", resp["image_byte_data"])
	}
}

func TestCreateImage(t *testing.T) {
	initDB()
	tx.SavePoint("sp1")

	// Request Body
	jsonBody := []byte(`{
		"id": 4,
		"ByteData": "hfji3u42987&(*&$#@(*Q&$ilkasmlkamljdaljo"
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	// Send new request whfji3u42987&(*&$#@(*Q&$ilkasmlkamljdaljoith json body info
	req, err := http.NewRequest("POST", "/images", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createImage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &resp)

	const mockData = "hfji3u42987&(*&$#@(*Q&$ilkasmlkamljdaljo"
	if resp["ByteData"] != mockData {
		t.Errorf("Image title is invalid, expected \"hfji3u42987&(*&$#@(*Q&$ilkasmlkamljdaljo\", got %v", resp["ByteData"])
	}

	tx.RollbackTo("sp1")
}

func TestUpdateImage(t *testing.T) {
	initDB()
	tx.SavePoint("sp2")

	// Request Body
	jsonBody := []byte(`{
		"id": 2,
		"ByteData": "dskjfo34($&*378uwefoisdjfoiwqurifj"
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	// Send new request with json body info
	req, err := http.NewRequest("PUT", "/images/2", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/images/{id}", updateImage)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &resp)

	const mockData = "dskjfo34($&*378uwefoisdjfoiwqurifj"
	if resp["ByteData"] != mockData {
		t.Errorf("Image Title is invalid, expected \"dskjfo34($&*378uwefoisdjfoiwqurifj\", got %v", resp["title"])
	}

	tx.RollbackTo("sp2")
}

func TestDeleteImage(t *testing.T) {
	initDB()
	tx.SavePoint("sp3")

	// Send new request with json body info
	req, err := http.NewRequest("DELETE", "/images/3", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/images/{id}", deleteImage)
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
