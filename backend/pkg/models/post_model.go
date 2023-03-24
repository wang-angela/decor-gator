package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID            uint    `gorm:"primaryKey" json:"id"`
	Title         string  `json:"title"`
	FurnitureType string  `json:"furnitureType"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	UserPosted    string  `json:"userPosted"` // Change later to user object once we learn how to test that on Postman.
}
