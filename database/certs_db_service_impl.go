package database

import (
	"github.com/howstrongiam/backend/models"
	"github.com/sirupsen/logrus"
)

type CertificationDbServiceImpl struct{}

func (s CertificationDbServiceImpl) AddNewCertification(cert models.Certification) *models.Certification {
	result := GetDbConn().Create(&cert)
	if result.Error != nil {
		logrus.Errorf("Unable to save certificate, %s", result.Error)
		return nil
	}
	var addedCert models.Certification
	result = GetDbConn().First(&addedCert, "id = ?", cert.ID)
	if result.Error != nil {
		logrus.Errorf("Unable to get certificate, %s", result.Error)
		return nil
	}
	return &addedCert
}

func (s CertificationDbServiceImpl) GetAllCertifications() []models.Certification {
	var certs []models.Certification
	result := GetDbConn().Find(&certs)
	if result.Error != nil {
		logrus.Errorf("Unable to get all certification, %s", result.Error)
		return nil
	}
	return certs
}
