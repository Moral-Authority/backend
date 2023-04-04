package database

import "github.com/howstrongiam/backend/models"

type CertificationDbService interface {
	AddNewCertification(cert models.Certification) *models.Certification
	GetAllCertifications() []models.Certification
}
