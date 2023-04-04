package models

import "gorm.io/gorm"

type ProductCategories struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	CategoryId uint `gorm:"foreignKey:CategoryRefer"`
	ProductId  uint `gorm:"foreignKey:ProductRefer"`
}

type Categories struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Name     string
	ParentId uint `gorm:"foreignKey:CategoryRefer"`
}
