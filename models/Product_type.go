package models

import "gorm.io/gorm"

type ProductType struct {
	gorm.Model
	ProductId uint `json:"product_id"`
}
