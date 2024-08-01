package models

import "gorm.io/gorm"

type Favourite struct {
	gorm.Model
	UserRefer uint
	User      User   `gorm:"foreignKey:UserRefer"`
	ProductId uint
	Product   Product `gorm:"foreignKey:ProductId"`
}