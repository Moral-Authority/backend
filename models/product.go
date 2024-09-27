package models

import "gorm.io/gorm"

type ProductBase struct {
	gorm.Model
	SubDepartment         int                    `gorm:"column:sub_department"`              // Maps to the "sub_department" column
	Category              string                 `gorm:"column:category"`                    // Maps to the "category" column
	Title                 string                 `gorm:"column:title"`                       // Maps to the "title" column
	Description           string                 `gorm:"column:description"`                 // Maps to the "description" column
	Url                   string                 `gorm:"column:url"`                         // Maps to the "url" column
	ProductImage          string                 `gorm:"column:product_image"`               // Maps to the "product_image" column
	CompanyID             uint                   `gorm:"index;column:company_id"`            // Index and maps to "company_id" column
	Company               Company                `gorm:"foreignKey:CompanyID;references:ID"` // Foreign key referencing "CompanyID" in the "Company" table
	ProductCertifications []ProductCertification `gorm:"foreignKey:ProductID"`               // Foreign key referencing "ProductID" in the "ProductCertification" table
	PurchaseInfo          []PurchaseInfo         `gorm:"foreignKey:ProductID"`               // Foreign key referencing "ProductID" in the "PurchaseInfo" table
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
	Price             float64
	Url               string
}
