package database

import "github.com/howstrongiam/backend/models"

type CertificatesDbService interface {
	AddNewCertificate(cert models.Certification) *models.Certification
}
