// Accessors used by internal func without mux router
package utils

import (
	"context"
	"log"

	"github.com/decor-gator/backend/pkg/configs"
	"github.com/decor-gator/backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var err error

// Used by CreatedTokenEndpoint() to check if users have an account
func JwtVerifyUserExists(user models.User) bool {
	var res models.User

	filter := bson.D{{Key: "username", Value: user.Username}}

	err := configs.GetCollection(configs.DB, "users").FindOne(context.TODO(), filter).Decode(&res)
	if err == mongo.ErrNoDocuments {
		return false
	} else if err != nil {
		// Throws error for other cases
		log.Fatal(err)
	}

	return true
}

func JwtVerifyPassword(user models.User) bool {
	var res models.User

	filter := bson.D{{Key: "username", Value: user.Username}}

	configs.GetCollection(configs.DB, "users").FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		// Decoding error, previous tests are done to ensure user exists
		log.Println("Error Decoding")
	}

	return ComparePassword(user.Password, res.Password)
}
