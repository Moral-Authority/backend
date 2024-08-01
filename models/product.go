package models

import "gorm.io/gorm"


type Product struct {
	gorm.Model
	Url         string
	Description string
	Title       string
}