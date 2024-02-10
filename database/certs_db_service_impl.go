package database

import (
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
)

type CertificationDbServiceImpl struct{}

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
		return nil, result.Error
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

func (s CertificationDbServiceImpl) UpdateCertification(cert models.Certification) *models.Certification {
	//cert := dbService.UpdateCertification(cert)
	//if cert == nil {
	//	return nil, errors.New("unable to get cert from db")
	//}
	//return toCertificationResponse(*cert), nil

	return &models.Certification{}
}
