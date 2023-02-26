package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ImageID       int    `gorm:"primaryKey" json:"id"`
	ImageByteData string `json:"byteData"`
}
