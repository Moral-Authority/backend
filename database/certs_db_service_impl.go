package database

import (
	"github.com/howstrongiam/backend/models"
	"github.com/sirupsen/logrus"
)

type CertificatesDbServiceImpl struct{}

func (s CertificatesDbServiceImpl) AddNewCertificate(cert models.Certification) *models.Certification {
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
