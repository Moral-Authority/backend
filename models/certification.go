package models

import (
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
)

type Certification struct {
	gorm.Model
	Name               null.String `json:"name"`
	Website            null.String `json:"website"`
	Logo               null.String `json:"logo"`
	Description        null.String `json:"description"`
	Industry           null.String `json:"industry"`
	Certifier          null.String `json:"certifying_company"`
	CertifiesCompany   null.Bool   `json:"certifies_company"`
	CertifiesProduct   null.Bool   `json:"certifies_product"`
	CertifiesProcess   null.Bool   `json:"certifies_process"`
	CertifierContactID null.String `json:"certifier_contact_id"`
	Audited            null.Bool   `json:"audited"`
	Auditor            null.String `json:"auditor"`
	Region             null.String `json:"region"`
	Qualifiers         null.String `json:"qualifiers"`
	Sources            null.String `json:"sources"`
	CompanyCertifications []CompanyCertification `gorm:"foreignKey:CertificationID"`
	ProductCertifications []ProductCertification `gorm:"foreignKey:CertificationID"`
}
