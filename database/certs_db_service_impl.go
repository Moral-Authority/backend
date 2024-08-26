package database

import (
	"fmt"

	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
)

type CertificationDbServiceImpl struct{}


func (s CertificationDbServiceImpl) GetCertificationsByFilter(filters map[string]interface{}, input model.FilterCertificationsInput) ([]models.Certification, int64, error) {
	var certs []models.Certification
	db := GetDbConn()

	// Apply filters to the query
	query := ApplyFilters(db, filters)

	// If pagination is provided, apply limit and offset
	if input.Pagination != nil {
		p := *input.Pagination
		offset := models.CalculateOffset(p)
		query = query.Limit(*input.Pagination.Items).Offset(offset)
	}

	// Execute the query and retrieve results
	if err := query.Find(&certs).Error; err != nil {
		logrus.Errorf("Unable to get certifications by filter, %s", err)
		return nil, 0, err
	}

	// Count the total number of records without pagination for metadata
	var total int64
	if err := db.Model(&models.Certification{}).Where(filters).Count(&total).Error; err != nil {
		logrus.Errorf("Unable to count certifications, %s", err)
		return nil, 0, err
	}

	return certs, total, nil
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
