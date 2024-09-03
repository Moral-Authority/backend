package models

import "gorm.io/gorm"

type ProductBase struct {
	gorm.Model
	SubDepartment         string
	Category              string
	Title                 string
	Description           string
	Url                   string
	ProductImage          string
	CompanyID             uint                   `gorm:"index"`
	Company               Company                `gorm:"foreignKey:CompanyID"`
	ProductCertifications []ProductCertification `gorm:"foreignKey:ProductID"`
	PurchaseInfo          []PurchaseInfo         `gorm:"foreignKey:ProductID"`
}

type HomeGardenProduct struct {
	ProductBase
}

type ClothingAccessoriesProduct struct {
	ProductBase
}

type HealthBathBeautyProduct struct {
	ProductBase
}

type ToysKidsBabiesProduct struct {
	ProductBase
}

type PurchaseInfo struct {
	gorm.Model
	ProductID         uint `gorm:"index"` // ID from the product table
	ProductDepartment int  `gorm:"index"` // Indicates which product table this ID belongs to
	Website           string
	Price             string
	Url               string
}
