package models

import (
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
)

type CompanyCertification struct {
	gorm.Model
	CompanyID       uint        `gorm:"foreignKey:CompanyRefer"`
	CertificationID uint        `gorm:"foreignKey:CertificationRefer"`
	CertifiedAt     null.Time   `json:"certified_at"`
	ExpirationDate  null.Time   `json:"expiration_date"`
	OtherDetails    null.String `json:"other_details"`
}
