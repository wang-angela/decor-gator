package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ID       int    `gorm:"primaryKey" json:"id"`
	ImageURL string `json:"url"`
}
