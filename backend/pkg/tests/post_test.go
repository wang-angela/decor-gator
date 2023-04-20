package tests

import (
	"bytes"
	"context"
	"encoding/json"
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

func TestGetAllPosts(t *testing.T) {

	// Send new request with json body info
	req, err := http.NewRequest("POST", "/posts", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetPosts)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp []models.Post
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Errorf("Invalid response, expected list of posts, got %v", rr.Body.String())
	}

	if len(resp) < 1 {
		t.Errorf("Invalid number of posts, expected 1, got %v", len(resp))
	}
}

func TestGetPost(t *testing.T) {
	configs.ConnectDB()

	// Send new request with json body info
	req, err := http.NewRequest("GET", "/posts/643dd5258fa45fce76227bbd", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/posts/{_id}", controllers.GetPost)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp *mongo.SingleResult
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp != nil {
		t.Errorf("Post not found")
	}
}

func TestCreatePost(t *testing.T) {
	configs.ConnectDB()
	controllers.InitAWSSession()
	controllers.CreateBucket()

	// Request Body
	jsonBody := []byte(`{
		"title":		 "A chair!",
		"furnitureType": "chair",
		"description":	 "such beautiful chair",
		"price": 		 22.23,
		"userPosted":	 "angela",
		"imageURL": 	 "fnewkdnfkkcdsmkfcwesd"
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	// Send new request with json body info
	req, err := http.NewRequest("POST", "/posts", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreatePost)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp *mongo.InsertOneResult
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp == nil {
		t.Errorf("Insert failed")
	}

	configs.GetCollection(configs.DB, "posts").DeleteOne(context.TODO(),
		bson.D{{Key: "_id", Value: resp.InsertedID}},
	)
}

func TestUpdatePost(t *testing.T) {
	// Request Body
	jsonBody := []byte(`{
		"title": "The Most Beautiful Desk",
		"description": "BUY ME!!!!!"
	}`)
	bodyReader := bytes.NewReader(jsonBody)

	// Send new request with json body info
	req, err := http.NewRequest("PUT", "/posts/643f3e7a30eabed5ba74768b", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/posts/{_id}", controllers.UpdatePost)
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decoding recorded response
	var resp *mongo.UpdateResult
	json.Unmarshal(rr.Body.Bytes(), &resp)

	if resp.MatchedCount == 0 {
		t.Errorf("Fail to update")
	}

	// Reversing changes
	update := bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "title", Value: "Beautiful Desk"},
			{Key: "furniture_type", Value: "Desk"},
			{Key: "description", Value: "What a nice desk."},
			{Key: "price", Value: 34.00},
			{Key: "user_posted", Value: "test@gmail.com"},
			{Key: "image_url", Value: "data:image/jpeg;base64,/9j/4QAYRXhpZgAASUkqAAgAAAAAAAAAAAAAAP/sABFEdWN"},
		},
	}}

	objectId, err := primitive.ObjectIDFromHex("643f3e7a30eabed5ba74768b")

	configs.GetCollection(configs.DB, "users").UpdateOne(context.TODO(),
		bson.D{{Key: "_id", Value: objectId}},
		update,
	)
}

func TestDeletePost(t *testing.T) {
	controllers.InitAWSSession()
	controllers.CreateBucket()

	post := &models.Post{
		ID:            primitive.NewObjectID(),
		Title:         "New Chair",
		FurnitureType: "Chair",
		Description:   "a new chair",
		Price:         10.0,
		UserPosted:    "jacksonSon",
		ImageURL:      "aURL",
	}

	configs.GetCollection(configs.DB, "posts").InsertOne(context.TODO(), post)

	// Send new request with json body info

	route := "/posts/"
	hexString := post.ID.Hex()
	route += hexString

	req, err := http.NewRequest("DELETE", route, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record Response
	rr := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/posts/{_id}", controllers.DeleteUser)
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
