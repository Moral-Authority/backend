package models

import "gorm.io/gorm"

type Certification struct {
	gorm.Model
	CertifyingCompany string `json:"certifying_company"`
	CertName          string `json:"cert_name"`
}
