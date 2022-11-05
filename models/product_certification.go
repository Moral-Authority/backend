package models

import "gorm.io/gorm"

type ProductCertification struct {
	gorm.Model
	CertificationId uint `json:"certification_id"`
	ProductId       uint `json:"product_id"`
}
