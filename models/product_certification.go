package models

import (
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
)

type ProductCertification struct {
	gorm.Model
	ProductID       uint          `gorm:"index"` // Foreign key referencing the product
	CertificationID uint          `gorm:"index"`
	Certification   Certification `gorm:"foreignKey:CertificationID"`
	CertifiedAt     null.Time     `json:"certified_at"`
	ExpirationDate  null.Time     `json:"expiration_date"`
	OtherDetails    null.String   `json:"other_details"`
}
