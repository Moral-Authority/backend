package models

import "gorm.io/gorm"

type Section struct {
	gorm.Model
	Title string `json:"title"`
}
