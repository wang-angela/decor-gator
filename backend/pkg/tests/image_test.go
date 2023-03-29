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

func TestGetAllImages(t *testing.T) {
	utils.InitDBTest("test")

	// Send new request with json body info
	req, err := http.NewRequest("GET", "/images", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetImages)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp []models.Image
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Errorf("Invalid response, expected list of images, got %v", rr.Body.String())
	}

	if len(resp) < 1 {
		t.Errorf("Invalid number of images, expected 1, got %v", len(resp))
	}
}

func TestGetImage(t *testing.T) {
	utils.InitDBTest("test")

	// Send new request with json body info
	req, err := http.NewRequest("GET", "/images/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/images/{id}", controllers.GetImage)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &resp)

	const url = "cool-url.go/img1"
	if resp["url"] != url {
		t.Errorf("Data is invalid, expected \"cool-url.go/img1\" got %v", resp["url"])
	}
}

func TestCreateImage(t *testing.T) {
	TX := utils.InitDBTest("test")
	TX.SavePoint("sp1")

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
	handler := http.HandlerFunc(controllers.CreateImage)

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

	TX.RollbackTo("sp1")
}

func TestUpdateImage(t *testing.T) {
	TX := utils.InitDBTest("test")
	TX.SavePoint("sp2")

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
	r.HandleFunc("/images/{id}", controllers.UpdateImage)
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

	TX.RollbackTo("sp2")
}

func TestDeleteImage(t *testing.T) {
	TX := utils.InitDBTest("test")
	TX.SavePoint("sp3")

	// Send new request with json body info
	req, err := http.NewRequest("DELETE", "/images/3", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/images/{id}", controllers.DeleteImage)
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
