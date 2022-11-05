package models

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	ProductId uint `json:"product_id"`
}
