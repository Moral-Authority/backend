package database

import (
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
	"fmt"
)

type CertificationDbServiceImpl struct{}

func (s CertificationDbServiceImpl) GetCertificationsByFilter(filters map[string]interface{}) ([]models.Certification, error) {
    var certs []models.Certification
    db := GetDbConn()

	query := ApplyFilters(db, filters)
	
    if err := query.Find(&certs).Error; err != nil {
        logrus.Errorf("Unable to get certifications by filter, %s", err)
        return nil, err
    }

    return certs, nil
}


func (s CertificationDbServiceImpl) AddNewCertification(cert models.Certification) (*models.Certification, error) {
	result := GetDbConn().Create(&cert)
	if result.Error != nil {
		logrus.Errorf("Unable to save certificate, %s", result.Error)
		return &cert, result.Error
	}
	var addedCert models.Certification
	result = GetDbConn().First(&addedCert, "id = ?", cert.ID)
	if result.Error != nil {
		logrus.Errorf("Unable to get certificate, %s", result.Error)
		return &cert, result.Error
	}
	return &addedCert, nil
}

func (s CertificationDbServiceImpl) GetAllCertifications() ([]models.Certification, error) {
	var certs []models.Certification
	result := GetDbConn().Find(&certs)
	if result.Error != nil {
		logrus.Errorf("Unable to get all certification, %s", result.Error)
		return nil, fmt.Errorf("failed to retrieve certifications: %w", result.Error)
	}
	return certs, nil
}

func (s CertificationDbServiceImpl) GetCertificationById(certId string) (*models.Certification, error) {
	var certs models.Certification
	result := GetDbConn().Find(&certs)
	if result.Error != nil {
		logrus.Errorf("Unable to get all certification, %s", result.Error)
		return &models.Certification{}, result.Error
	}
	return &certs, nil
}

func (s CertificationDbServiceImpl) UpdateCertification(cert models.Certification) (*models.Certification, error) {
    var existingCert models.Certification
    result := GetDbConn().First(&existingCert, "id = ?", cert.ID)
    if result.Error != nil {
        logrus.Errorf("Unable to find certification, %s", result.Error)
        return nil, result.Error
    }

    // Update the existing certification with the new data
    result = GetDbConn().Model(&existingCert).Updates(cert)
    if result.Error != nil {
        logrus.Errorf("Unable to update certification, %s", result.Error)
        return nil, result.Error
    }

    return &existingCert, nil
}
