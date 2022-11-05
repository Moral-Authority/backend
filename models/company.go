package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Url         string `json:"url"`
	Description string `json:"description"`
	UserId      uint   `json:"user_id"`
	IsVerified  bool   `json:"is_verified"`
	ImageId     uint   `json:"image_id"`
}
