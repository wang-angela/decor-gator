/* Code referenced from:
* https://www.youtube.com/watch?v=KPftgI40WHI by the Daily Code Buffer
* https://blog.logrocket.com/routing-go-gorilla-mux/ by Paul Akinyemi
 */

package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/decor-gator/backend/pkg/configs"
	"github.com/decor-gator/backend/pkg/models"
	"github.com/decor-gator/backend/pkg/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var err error
var userColl *mongo.Collection = configs.GetCollection(configs.DB, "users")

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var users []models.User

	// Retrieving users
	cur, err := userColl.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	// Decodes results into User array
	if cur.All(context.TODO(), &users) != nil {
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(&users)
	if err != nil {
		log.Printf("Error Encoding.")
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var user models.User

	// Prints an error id the user doesn't exists.
	params := mux.Vars(r)["email"]
	filter := bson.D{{Key: "email", Value: params}}

	err := userColl.FindOne(context.TODO(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		// Throws error if no user exists
		msg := "User does not exist"

		err = json.NewEncoder(w).Encode(&msg)
		if err != nil {
			log.Fatalln("Error Encoding")
		}
		return
	} else if err != nil {
		// Throws error for other cases
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Returns error if decoding is unsuccessful
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	// Encrypting password and generating id
	user.Password = utils.Encrypt(user.Password)
	user.ID = primitive.NewObjectID()

	// Insert into MongoDB
	res, err := userColl.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	email := []string{user.Email}
	SendWelcomeEmail(email)

	// Returns error if encoding is unsuccessful
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	// Update password
	user.Password = utils.Encrypt(user.Password)

	// Prints an error id the user doesn't exists.
	params := mux.Vars(r)["email"]
	filter := bson.D{{Key: "email", Value: params}}

	update := bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "first_name", Value: user.FirstName},
			{Key: "last_name", Value: user.LastName},
			{Key: "email", Value: user.Email},
			{Key: "username", Value: user.Username},
			{Key: "password", Value: user.Password},
		},
	}}

	res, err := userColl.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	} else if res.MatchedCount == 0 {
		msg := "User does not exist"

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

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")

	// Search parameters
	params := mux.Vars(r)["email"]
	filter := bson.D{{Key: "email", Value: params}}

	res, err := userColl.DeleteOne(context.TODO(), filter)
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
