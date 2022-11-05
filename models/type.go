package models

import "gorm.io/gorm"

type Type struct {
	gorm.Model
	Title      string `json:"title"`
	CategoryId uint   `json:"category_id"`
}
