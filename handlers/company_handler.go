package handlers

import (
	"errors"
	"fmt"

	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"github.com/volatiletech/null/v8"
)

type CompanyService struct{}

func (s CompanyService) GetCompaniesByFilter(dbService database.CompanyDbService, filters map[string]interface{}) ([]*model.Company, error) {
	companies, err := dbService.GetCompaniesByFilter(filters)
	if err != nil {
		return nil, err
	}

	var result []*model.Company
	for _, c := range companies {
		result = append(result, toCompanyResponse(&c))
	}

	return result, nil
}

func (s CompanyService) AddCompany(request model.AddCompany, dbService database.CompanyDbService, imageDbService database.ImageDbService, certDbService database.CertificationDbService) (*model.Company, error) {
	// _, err := database.StringToUint(request.UserID)
	// if err != nil {
	//     return nil, err
	// }

	company := models.Company{
		Name: request.Name,
	}

	if request.Logo != nil {
		company.Image = null.StringFrom(*request.Logo)
	}

	if request.URL != nil {
		company.Url = null.StringFrom(*request.URL)

	}

	if request.Description != nil {
		company.Description = null.StringFrom(*request.Description)

	}

	if request.City != nil {
		company.City = null.StringFrom(*request.City)

	}

	if request.State != nil {
		company.State = null.StringFrom(*request.State)

	}

	if request.Country != nil {
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

	savedCompany, err := dbService.AddCompany(company)
	if err != nil || savedCompany == nil {
		return nil, errors.New("unable to save company in db")
	}
	return toCompanyResponse(savedCompany), nil
}

func (s CompanyService) GetCompanyById(companyId string, dbService database.CompanyDbService) (*model.Company, error) {
	company, err := dbService.GetCompanyById(companyId)
	if err != nil || company == nil {
		return nil, errors.New("unable to get company from db")
	}
	return toCompanyResponse(company), nil
}

func (s CompanyService) GetAllCompanies(dbService database.CompanyDbService) ([]*model.Company, error) {
	companies, err := dbService.GetAllCompanies()
	if err != nil || companies == nil {
		return nil, errors.New("unable to get companies from db")
	}

	return toCompaniesResponse(companies), nil
}

func (s CompanyService) UpdateCompany(dbService database.CompanyDbService, input model.UpdateCompany) (*model.Company, error) {
	// Retrieve the existing certification
	currentCompany, err := dbService.GetCompanyById(input.ID)
	if err != nil || currentCompany == nil {
		return nil, fmt.Errorf(fmt.Sprintf("unable to get cert from db, ERROR: %s", err))
	}

	if input.Name != nil {
		currentCompany.Name = *input.Name
	}

	if input.Description != nil {
		currentCompany.Description = null.StringFrom(*input.Description)
	}

	if input.URL != nil {
		currentCompany.Url = null.StringFrom(*input.URL)
	}

	if input.Logo != nil {
		currentCompany.Image = null.StringFrom(*input.Logo)
	}

	if input.City != nil {
		currentCompany.City = null.StringFrom(*input.City)
	}

	if input.State != nil {
		currentCompany.State = null.StringFrom(*input.State)
	}

	if input.Country != nil {
		currentCompany.Country = null.StringFrom(*input.Country)
	}

	if input.IsVerified != nil {
		currentCompany.IsVerified = null.BoolFrom(*input.IsVerified)
	}

	//  TODO update certifications
	updatedCompany, err := dbService.UpdateCompany(*currentCompany)
	if err != nil || updatedCompany == nil {
		return nil, err
	}
	return toCompanyResponse(updatedCompany), nil
}
