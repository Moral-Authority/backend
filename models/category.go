package models

import "gorm.io/gorm"

type ProductCategories struct {
	gorm.Model
	CategoryId uint `gorm:"foreignKey:CategoryRefer"`
	ProductId  uint `gorm:"foreignKey:ProductRefer"`
}

type Category struct {
	gorm.Model
	Name     string
	Type     string
	ParentId *uint // `gorm:"foreignKey:CategoryRefer"`
}
