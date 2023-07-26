package handlers

import (
	"errors"
	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/models"
	"github.com/volatiletech/null/v8"
)

type CompanyService struct{}

func (s CompanyService) AddCompany(request model.AddCompany, dbService database.CompanyDbService, imageDbService database.ImageDbService, certDbService database.CertificationDbService) (*model.Company, error) {
	//userId, err := database.StringToUint(request.UserID)
	//if err != nil {
	//	return nil, err
	//}

	company := models.Company{
		Name: *request.Name,
	}

	if &request.Logo != nil {
		company.Image = null.StringFrom(*request.Logo)
	}

	if &request.URL != nil {
		company.Url = null.StringFrom(*request.URL)

	}

	if &request.Description != nil {
		company.Description = null.StringFrom(*request.Description)

	}

	if &request.City != nil {
		company.City = null.StringFrom(*request.City)

	}

	if &request.State != nil {
		company.State = null.StringFrom(*request.State)

	}

	if &request.Country != nil {
		company.Country = null.StringFrom(*request.Country)

	}

	//cert := models.Certification{
	//	CertifyingCompany: "request.Certification.CertifyingCompany",
	//	CertName:          "request.Certification.CertName",
	//}
	//
	//addedCert := certDbService.AddNewCertification(cert)
	//if addedCert == nil {
	//	return nil, errors.New("unable to save certificate to db")
	//}

	savedCompany := dbService.AddCompany(company)
	if savedCompany == nil {
		return nil, errors.New("unable to save company in db")
	}
	return toCompanyResponse(savedCompany), nil
	return nil, nil
}

func (s CompanyService) GetCompanyById(companyId string, dbService database.CompanyDbService) (*model.Company, error) {
	company := dbService.GetCompanyById(companyId)
	if company == nil {
		return nil, errors.New("unable to get company from db")
	}
	return toCompanyResponse(company), nil
}

func (s CompanyService) GetAllCompanies(dbService database.CompanyDbService) ([]*model.Company, error) {
	companies := dbService.GetAllCompanies()
	if companies == nil {
		return nil, errors.New("unable to get companies from db")
	}
	return toCompaniesResponse(companies), nil
}
