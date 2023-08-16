package database

import (
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
)

type CertificationDbServiceImpl struct{}

func (s CertificationDbServiceImpl) AddNewCertification(cert models.Certification) *models.Certification {
	result := GetDbConn().Create(&cert)
	if result.Error != nil {
		logrus.Errorf("Unable to save certificate, %s", result.Error)
		return &cert
	}
	var addedCert models.Certification
	result = GetDbConn().First(&addedCert, "id = ?", cert.ID)
	if result.Error != nil {
		logrus.Errorf("Unable to get certificate, %s", result.Error)
		return &cert
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

func (s CertificationDbServiceImpl) GetCertificationById(certId string) *models.Certification {
	var certs models.Certification
	result := GetDbConn().Find(&certs)
	if result.Error != nil {
		logrus.Errorf("Unable to get all certification, %s", result.Error)
		return &models.Certification{}
	}
	return &certs
}

func (s CertificationDbServiceImpl) UpdateCertification(cert models.Certification) *models.Certification {
	//cert := dbService.UpdateCertification(cert)
	//if cert == nil {
	//	return nil, errors.New("unable to get cert from db")
	//}
	//return toCertificationResponse(*cert), nil

	return &models.Certification{}
}
