package handlers

import (
	"errors"
	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"github.com/volatiletech/null/v8"
)

type CertificationService struct{}

func (s CertificationService) AddNewCertification(request model.AddCertification, dbService database.CertificationDbService) (*model.Certification, error) {

    cert := models.Certification{
        Name:               null.StringFrom(*request.Name),
        Logo:               null.StringFrom(*request.Logo),
        Industry:           null.StringFrom(*request.Industry),
        Certifier:          null.StringFrom(*request.Certifier),
        CertifiesCompany: null.BoolFrom(*request.CertifiesCompany),
        CertifiesProduct:  null.BoolFrom(*request.CertifiesProduct),
        CertifiesProcess: null.BoolFrom(*request.CertifiesProcess),
        CertifierContactID: null.StringFrom(*request.CertifierContactID), 
        Audited:            null.BoolFrom(*request.Audited),
        Auditor:            null.StringFrom(*request.Auditor),
        Region:             null.StringFrom(*request.Region),
        Qualifiers:         null.StringFrom(*request.Qualifiers),
        Sources:            null.StringFrom(*request.Sources),
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

func (s CertificationService) GetCertificationById(certId string, dbService database.CertificationDbService) (*model.Certification, error) {
	cert := dbService.GetCertificationById(certId)
	if cert == nil {
		return nil, errors.New("unable to get cert from db")
	}
	return toCertificationResponse(*cert), nil
}

func (s CertificationService) UpdateCertification(request model.UpdateCertification, dbService database.CertificationDbService) (*model.Certification, error) {
    // Retrieve the existing certification
    cert := dbService.GetCertificationById(request.ID)
    if cert == nil {
        return nil, errors.New("unable to find certificate in db")
    }

    // Update the fields
    cert.Name = null.StringFrom(*request.Name)
    cert.Logo = null.StringFrom(*request.Logo)
    cert.Industry = null.StringFrom(*request.Industry)
    cert.Certifier = null.StringFrom(*request.Certifier)
    cert.CertifiesCompany = null.BoolFrom(*request.CertifiesCompany)
    cert.CertifiesProduct = null.BoolFrom(*request.CertifiesProduct)
    // cert.CertifiesProcess = null.BoolFrom(*request.CertifiesProcess)
    cert.CertifierContactID = null.StringFrom(*request.CertifierContactID)
    cert.Audited = null.BoolFrom(*request.Audited)
    cert.Auditor = null.StringFrom(*request.Auditor)
    cert.Region = null.StringFrom(*request.Region)
    cert.Qualifiers = null.StringFrom(*request.Qualifiers)
    cert.Sources = null.StringFrom(*request.Sources)

    // Save the updated certification
    updatedCert := dbService.UpdateCertification(*cert)
    if updatedCert == nil {
        return nil, errors.New("unable to update certificate in db")
    }

    return toCertificationResponse(*updatedCert), nil
}