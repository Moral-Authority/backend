package models

import "gorm.io/gorm"

type ProductCategories struct {
	gorm.Model
	CategoryID uint `gorm:"foreignKey:CategoryRefer"`
	ProductID  uint `gorm:"foreignKey:ProductRefer"`
}

type Category struct {
	gorm.Model
	Name     string
	Type     string
	ParentID *uint `gorm:"foreignKey:CategoryRefer"`
	Products []Product `gorm:"many2many:product_categories;"` // Many-to-many relationship with Product
}
