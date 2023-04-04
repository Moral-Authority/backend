package handlers

import (
	"errors"
	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/models"
	"github.com/volatiletech/null/v8"
)

type CertificationService struct{}

func (s CertificationService) AddNewCertification(request model.AddCertification, dbService database.CertificationDbService) (*model.Certification, error) {

	cert := models.Certification{
		Name:             null.StringFrom(*request.Name),
		Logo:             null.StringFrom(*request.Logo),
		Industry:         null.StringFrom(*request.Industry),
		Certifier:        null.StringFrom(*request.Certifier),
		CertifiesCompany: null.BoolFrom(*request.CertifiesCompany),
		CertifiesProduct: null.BoolFrom(*request.CertifiesProduct),
		CertifiesProcess: null.BoolFrom(*request.CertifiesProcess),
		CertifierContact: null.StringFrom(*request.CertifierContact),
		Audited:          null.BoolFrom(*request.Audited),
		Auditor:          null.StringFrom(*request.Auditor),
		Region:           null.StringFrom(*request.Region),
		Qualifiers:       null.StringFrom(*request.Qualifiers),
		Sources:          null.StringFrom(*request.Sources),
	}

	addedCert := dbService.AddNewCertification(cert)
	if addedCert == nil {
		return nil, errors.New("unable to save certificate to db")
	}

	return toCertificationResponse(*addedCert), nil
}

func (s CertificationService) GetAllCertifications(dbService database.CertificationDbService) ([]*model.Certification, error) {
	certs := dbService.GetAllCertifications()
	var response []*model.Certification
	for _, e := range certs {
		cert := toCertificationResponse(e)
		response = append(response, cert)
	}
	return response, nil
}

//func (s CompanyService) GetCertificationById(certId string, dbService database.CertificationDbService) (*model.Certification, error) {
//	cert := dbService.GetCertificationById(certId)
//	if cert == nil {
//		return nil, errors.New("unable to get cert from db")
//	}
//	return toCertificationResponse(*cert), nil
//}
