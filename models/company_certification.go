package models

import "gorm.io/gorm"

type CompanyCertification struct {
	gorm.Model
	CertificationId uint `json:"certification_id"`
	CompanyId       uint `json:"company_id"`
}
