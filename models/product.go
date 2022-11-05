package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Url         string `json:"url"`
	Description string `json:"description"`
	UserId      uint   `json:"user_id"`
	ImageId     uint   `json:"image_id"`
}
