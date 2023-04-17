package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	Title         string             `json:"title"`
	FurnitureType string             `json:"furnitureType"`
	Description   string             `json:"description"`
	Price         float64            `json:"price"`
	UserPosted    string             `json:"userPosted"` // Change later to user object once we learn how to test that on Postman.
	ImageURL      string             `json:"imageURL"`
}
