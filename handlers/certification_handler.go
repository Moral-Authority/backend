package handlers

import (
	"fmt"

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

func (s CertificationService) GetAllCertifications(dbService database.CertificationDbService) ([]*model.Certification, error) {
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

func (s CertificationService) GetCertificationById(certId string, dbService database.CertificationDbService) (*model.Certification, error) {
	cert, err := dbService.GetCertificationById(certId)
	if err != nil || cert == nil {
		return nil, fmt.Errorf("unable to get cert from db, ERROR: %s", err)
	}

	return toCertificationResponse(*cert), nil
}

func (s CertificationService) UpdateCertification(request model.UpdateCertification, dbService database.CertificationDbService) (*model.Certification, error) {
	// Retrieve the existing certification
	cert, err := dbService.GetCertificationById(request.ID)
	if err != nil || cert == nil {
		return nil, fmt.Errorf(fmt.Sprintf("unable to get cert from db, ERROR: %s", err))
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
	updatedCert,err := dbService.UpdateCertification(*cert)
	if err != nil ||updatedCert == nil {
		return nil, err
	}

	return toCertificationResponse(*updatedCert), nil
}
