package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ID            int    `gorm:"primaryKey" json:"id"`
	ImageByteData string `json:"ByteData"`
}
