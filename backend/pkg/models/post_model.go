package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	Title         string             `json:"title" bson:"title"`
	FurnitureType string             `json:"furnitureType" bson:"furniture_type"`
	Description   string             `json:"description" bson:"description"`
	Price         float64            `json:"price" bson:"price"`
	UserPosted    string             `json:"userPosted" bson:"user_posted"` // Change later to user object once we learn how to test that on Postman.
	ImageURL      string             `json:"imageURL" bson:"image_url"`
}
