package models

import (
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
)

type CompanyCertification struct {
	gorm.Model
	CompanyID       uint          `gorm:"foreignKey:CompanyRefer"`
	Company         Company       `gorm:"foreignKey:CompanyID"`
	CertificationID uint          `gorm:"foreignKey:CertificationRefer"`
	Certification   Certification `gorm:"foreignKey:CertificationID"` // This field should be added
	CertifiedAt     null.Time     `json:"certified_at"`
	ExpirationDate  null.Time     `json:"expiration_date"`
	OtherDetails    null.String   `json:"other_details"`
}
