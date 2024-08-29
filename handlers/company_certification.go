package handlers

import (
	"errors"

	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
)

func (s CompanyService) AddCompanyCertificationHandler(companyService database.CompanyDbService, certService database.CertificationDbService, request model.CompanyCertificationInput) (*models.CompanyCertification, error) {

	
	company, err := companyService.GetCompanyById(request.CompanyID)
	if err != nil {
		return nil, errors.New("unable to get company from db")
	}

	cert, err := certService.GetCertificationById(request.CertificationID)
	if err != nil {
		return nil, errors.New("unable to get certification from db")
	}

	cc := models.CompanyCertification{}
	cc.Company.ID = company.ID
	cc.Certification.ID = cert.ID

	addedCert, err := companyService.AddCompanyCertification(cc)
	if err != nil {
		return nil, errors.New("unable to save company certification in db")
	}

	return addedCert, nil
}
