package models

import (
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
)

type CompanyCertification struct {
	gorm.Model
	CompanyID       uint          `gorm:"index"` // Index for performance optimization
	Company         Company       `gorm:"foreignKey:CompanyID"`
	CertificationID uint          `gorm:"index"`
	Certification   Certification `gorm:"foreignKey:CertificationID"`
	CertifiedAt     null.Time     `json:"certified_at"`
	ExpirationDate  null.Time     `json:"expiration_date"`
	OtherDetails    null.String   `json:"other_details"`
}