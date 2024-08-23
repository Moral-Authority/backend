package models

import (
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
)

type CompanyProduct struct {
    gorm.Model
    CompanyID      uint        `gorm:"column:company_id"` // Corrected foreign key tag
    ProductID      uint        `gorm:"column:product_id"` // Corrected foreign key tag
    LaunchedAt     null.Time   `json:"launched_at"`
    DiscontinuedAt null.Time   `json:"discontinued_at"`
    OtherDetails   null.String `json:"other_details"`
    Company        Company     `gorm:"foreignKey:CompanyID"`
    Product        Product     `gorm:"foreignKey:ProductID"`
}
