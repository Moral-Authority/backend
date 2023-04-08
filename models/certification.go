package models

import (
	"github.com/volatiletech/null/v8"
	"gorm.io/gorm"
)

type Certification struct {
	gorm.Model
	Name               null.String `json:"name"`
	Logo               null.String `json:"logo"`
	Industry           null.String `json:"industry"`
	Certifier          null.String `json:"certifying_company"`
	CertifiesCompanies null.Bool   `json:"certifies_company"`
	CertifiesProducts  null.Bool   `json:"certifies_product"`
	CertifiesProcesses null.Bool   `json:"certifies_process"`
	CertifierContactID null.String `json:"certifier_contact_id"`
	Audited            null.Bool   `json:"audited"`
	Auditor            null.String `json:"auditor"`
	Region             null.String `json:"region"`
	Qualifiers         null.String `json:"qualifiers"`
	Sources            null.String `json:"sources"`
}
