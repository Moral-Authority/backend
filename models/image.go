package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	Url string
	ProductID  uint `gorm:"foreignKey:ProductRefer"`
}