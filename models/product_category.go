package models

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	ProductId uint `json:"product_id"`
}
