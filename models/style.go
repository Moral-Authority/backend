package models

import "gorm.io/gorm"

type Style struct {
	gorm.Model
	Title  string `json:"title"`
	TypeId uint   `json:"type_id"`
}
