package database

import "github.com/howstrongiam/backend/models"

type CompanyDbService interface {
	AddCompany(company models.Company) *models.Company
	GetCompanyById(companyId string) *models.Company
}
