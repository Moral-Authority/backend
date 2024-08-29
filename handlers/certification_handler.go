package handlers

import (
	"fmt"

	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"github.com/volatiletech/null/v8"
)

type CertificationService struct{}

func (s CertificationService) GetCertificationsByFilterHandler(filters map[string]interface{}, input model.FilterCertificationsInput, dbService database.CertificationDbService) ([]*model.Certification, int64, error) {
    certs, total, err := dbService.GetCertificationsByFilter(filters, input)
    if err != nil {
        return nil, 0, err
    }

    var result []*model.Certification
    for _, cert := range certs {
        result = append(result, toCertificationResponse(cert))
    }

    return result, total, nil
}

func (s CertificationService) AddNewCertificationHandler(request model.AddCertification, dbService database.CertificationDbService) (*model.Certification, error) {

	cert := models.Certification{
		Name:               null.StringFrom(*request.Name),
		Logo:               null.StringFrom(*request.Logo),
		Industry:           null.StringFrom(*request.Industry),
		Certifier:          null.StringFrom(*request.Certifier),
		CertifiesCompany:   null.BoolFrom(*request.CertifiesCompany),
		CertifiesProduct:   null.BoolFrom(*request.CertifiesProduct),
		CertifiesProcess:   null.BoolFrom(*request.CertifiesProcess),
		CertifierContactID: null.StringFrom(*request.CertifierContactID),
		Audited:            null.BoolFrom(*request.Audited),
		Auditor:            null.StringFrom(*request.Auditor),
		Region:             null.StringFrom(*request.Region),
		Qualifiers:         null.StringFrom(*request.Qualifiers),
		Sources:            null.StringFrom(*request.Sources),
	}

	addedCert, err := dbService.AddNewCertification(cert)
	if err != nil || addedCert == nil {
		return nil, fmt.Errorf(fmt.Sprintf("unable to add cert to db, ERROR: %s", err))
	}

	return toCertificationResponse(*addedCert), nil
}

func (s CertificationService) GetAllCertificationsHandler(dbService database.CertificationDbService) ([]*model.Certification, error) {
	certs, err := dbService.GetAllCertifications()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("unable to get certs from db, ERROR: %s", err))
	}

	var response []*model.Certification
	for _, e := range certs {
		cert := toCertificationResponse(e)
		response = append(response, cert)
	}
	return response, nil
}

func (s CertificationService) GetCertificationByIdHandler(certId string, dbService database.CertificationDbService) (*model.Certification, error) {
	cert, err := dbService.GetCertificationById(certId)
	if err != nil || cert == nil {
		return nil, fmt.Errorf("unable to get cert from db, ERROR: %s", err)
	}

	return toCertificationResponse(*cert), nil
}

func (s CertificationService) UpdateCertificationHandler(request model.UpdateCertification, dbService database.CertificationDbService) (*model.Certification, error) {
	// Retrieve the existing certification
	cert, err := dbService.GetCertificationById(request.ID)
	if err != nil || cert == nil {
		return nil, fmt.Errorf(fmt.Sprintf("unable to get cert from db, ERROR: %s", err))
	}

	// Update the fields
	if request.Name != nil {
		cert.Name = null.StringFrom(*request.Name)
	}

	if request.Logo != nil {
		cert.Logo = null.StringFrom(*request.Logo)
	}

	if request.Industry != nil {
		cert.Industry = null.StringFrom(*request.Industry)
	}

	if request.Certifier != nil {
		cert.Certifier = null.StringFrom(*request.Certifier)
	}

	if request.CertifiesCompany != nil {
		cert.CertifiesCompany = null.BoolFrom(*request.CertifiesCompany)	
	}

	if request.CertifiesProduct != nil {
		cert.CertifiesProduct = null.BoolFrom(*request.CertifiesProduct)
	}

	// if request.CertifiesProcess != nil {
	// 	cert.CertifiesProcess = null.BoolFrom(*request.CertifiesProcess)
	// }

	if request.CertifierContactID != nil {
		cert.CertifierContactID = null.StringFrom(*request.CertifierContactID)
	}

	if request.Audited != nil {
		cert.Audited = null.BoolFrom(*request.Audited)
	}

	if request.Auditor != nil {
		cert.Auditor = null.StringFrom(*request.Auditor)
	}

	if request.Region != nil {		
		cert.Region = null.StringFrom(*request.Region)
	}

	if request.Qualifiers != nil {
		cert.Qualifiers = null.StringFrom(*request.Qualifiers)
	}

	if request.Sources != nil {
		cert.Sources = null.StringFrom(*request.Sources)
	}	

	// Save the updated certification
	updatedCert,err := dbService.UpdateCertification(*cert)
	if err != nil ||updatedCert == nil {
		return nil, err
	}

	return toCertificationResponse(*updatedCert), nil
}