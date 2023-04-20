package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/decor-gator/backend/pkg/configs"
	"github.com/decor-gator/backend/pkg/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var postColl *mongo.Collection = configs.GetCollection(configs.DB, "posts")

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var posts []models.Post

	// Retrieving posts
	cur, err := postColl.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	// Decodes results into Post array
	if cur.All(context.TODO(), &posts) != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(&posts)
	if err != nil {
		log.Printf("Error Encoding.")
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var post models.Post

	// Convert params string into ObjectID
	params := mux.Vars(r)["id"]
	objectId, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		log.Println("Invalid id")
	}

	// Search for posts, throw error if no document exists
	filter := bson.D{{Key: "_id", Value: objectId}}

	err = postColl.FindOne(context.TODO(), filter).Decode(&post)
	if err == mongo.ErrNoDocuments {
		// Throws error if no post exists
		msg := "Post does not exist"

		err = json.NewEncoder(w).Encode(&msg)
		if err != nil {
			log.Fatalln("Error Encoding")
		}
		return
	} else if err != nil {
		// Throws error for other cases
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(&post)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Returns error if decoding is unsuccessful
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	// Generating id
	post.ID = primitive.NewObjectID()

	// Insert into MongoDB
	res, err := postColl.InsertOne(context.TODO(), post)
	if err != nil {
		log.Fatal(err)
	}

	// Returns error if encoding is unsuccessful
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Fatalln("Error Encoding")
	}

	UploadObject(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	// Convert params string into ObjectID
	params := mux.Vars(r)["_id"]
	objectId, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		log.Println("Invalid id")
	}

	// Prints an error id the post doesn't exists.
	filter := bson.D{{Key: "_id", Value: objectId}}

	update := bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "title", Value: post.Title},
			{Key: "furniture_type", Value: post.FurnitureType},
			{Key: "description", Value: post.Description},
			{Key: "price", Value: post.Price},
			{Key: "user_posted", Value: post.UserPosted},
			{Key: "image_url", Value: post.ImageURL},
		},
	}}

	res, err := postColl.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	} else if res.MatchedCount == 0 {
		msg := "Post does not exist"

		err = json.NewEncoder(w).Encode(&msg)
		if err != nil {
			log.Fatalln("Error Encoding")
		}
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")

	// Convert params string into ObjectID
	params := mux.Vars(r)["id"]
	objectId, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		log.Println("Invalid id")
	}

	// Search parameters
	filter := bson.D{{Key: "_id", Value: objectId}}

	res, err := postColl.DeleteOne(context.TODO(), filter)
	if err != nil {
		// Throws error for other cases
		log.Fatal(err)
	} else if res.DeletedCount == 0 {
		msg := "Failed to delete"

		err = json.NewEncoder(w).Encode(&msg)
		if err != nil {
			log.Fatalln("Error Encoding")
		}
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Fatalln("Error Encoding")
	}

	DeleteObject(mux.Vars(r)["id"])
}
