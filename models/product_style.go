package models

import "gorm.io/gorm"

type ProductStyle struct {
	gorm.Model
	ProductId uint `json:"product_id"`
}
