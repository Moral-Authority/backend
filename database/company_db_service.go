package database

import (
	"github.com/Moral-Authority/backend/models"
)

type CompanyDbService interface {
	AddCompany(company models.Company) (*models.Company, error)
	GetCompanyById(companyId string) (*models.Company, error)
	GetAllCompanies() ([]*models.Company, error)
	GetCompaniesByFilter(filters map[string]interface{}) ([]models.Company, error)
	UpdateCompany(company models.Company) (*models.Company, error)
	AddCompanyCertification(request models.CompanyCertification) (*models.CompanyCertification, error)
	FindCompanyIDByName(companyName string) (uint, error)
}
