package models

import "gorm.io/gorm"

type Favourite struct {
	gorm.Model
	UserId    uint `json:"user_id"`
	ProductId uint `json:"product_id"`
}
