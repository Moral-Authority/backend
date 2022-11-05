package models

import "gorm.io/gorm"

type ProductDepartment struct {
	gorm.Model
	ProductId uint `json:"product_id"`
}
