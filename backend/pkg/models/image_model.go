package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ID  int    `gorm:"primaryKey" json:"id"`
	URL string `json:"url"`
}
