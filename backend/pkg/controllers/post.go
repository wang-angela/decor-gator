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

	// Prints an error id the post doesn't exists.
	params := mux.Vars(r)["id"]
	filter := bson.D{{Key: "id", Value: params}}

	err := postColl.FindOne(context.TODO(), filter).Decode(&post)
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
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	// Prints an error id the post doesn't exists.
	params := mux.Vars(r)["_id"]
	filter := bson.D{{Key: "_id", Value: params}}

	update := bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "_id", Value: post.ID},
			{Key: "title", Value: post.Title},
			{Key: "furnitureType", Value: post.FurnitureType},
			{Key: "description", Value: post.Description},
			{Key: "price", Value: post.Price},
			{Key: "userPosted", Value: post.UserPosted},
			{Key: "imageURL", Value: post.ImageURL},
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

	// Search parameters
	params := mux.Vars(r)["_id"]
	filter := bson.D{{Key: "_id", Value: params}}

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
}
