package database

import "github.com/Moral-Authority/backend/models"

type CertificationDbService interface {
    AddNewCertification(cert models.Certification) (*models.Certification, error)
	GetAllCertifications() ([]models.Certification, error)
	UpdateCertification(cert models.Certification) (*models.Certification, error)
    GetCertificationById(id string) (*models.Certification, error)
	GetCertificationsByFilter(filters map[string]interface{}) ([]models.Certification, error)
}
