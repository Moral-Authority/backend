package database

import "github.com/Moral-Authority/backend/models"

type CertificationDbService interface {
	AddNewCertification(cert models.Certification) *models.Certification
	GetAllCertifications() []models.Certification
	UpdateCertification(cert models.Certification) *models.Certification
	GetCertificationById(id string) *models.Certification
}
