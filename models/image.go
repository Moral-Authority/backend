package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ImageLocation string `json:"image_location"`
}
