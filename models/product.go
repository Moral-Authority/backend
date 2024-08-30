package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title                 string
	Description           string
	Url                   string
	CompanyID             uint                   `gorm:"index"` // Foreign key for Company
	Company               Company                // The related Company
	ProductCertifications []ProductCertification `gorm:"foreignKey:ProductID"`
	Images                []Image                `gorm:"foreignKey:ProductID"`
	PurchaseInfo          []PurchaseInfo         `gorm:"foreignKey:ProductID"`
}

type PurchaseInfo struct {
	gorm.Model
	ProductID uint   // GORM will automatically use this as a foreign key
	Website   string
	Price     string
	Url       string
}
