package database

import (
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
)

type CompanyDbServiceImpl struct{}

func (s CompanyDbServiceImpl) AddCompany(company models.Company) *models.Company {
	result := GetDbConn().Create(&company)
	if result.Error != nil {
		logrus.Errorf("Unable to save company, %s", result.Error)
		return nil
	}

	var addedCompany models.Company
	result = GetDbConn().First(&addedCompany, "id = ?", company.ID)
	if result.Error != nil {
		logrus.Errorf("Unable to get company, %s", result.Error)
		return nil
	}
	return &addedCompany
}

func (s CompanyDbServiceImpl) GetCompanyById(companyId string) *models.Company {
	var company models.Company
	result := GetDbConn().First(&company, "id = ?", companyId)
	if result.Error != nil {
		logrus.Errorf("Unable to get company, %s", result.Error)
		return nil
	}
	return &company
}

func (s CompanyDbServiceImpl) GetAllCompanies() []*models.Company {
	var companies []*models.Company
	//result := GetDbConn().First(&company, "id = ?", companyId)
	//if result.Error != nil {
	//	logrus.Errorf("Unable to get company, %s", result.Error)
	//	return nil, result.Error
	//}
	return companies
}
