package database

import (
    "github.com/Moral-Authority/backend/models"
    "github.com/sirupsen/logrus"
)

type CompanyDbServiceImpl struct{}

func (s CompanyDbServiceImpl) AddCompany(company models.Company) (*models.Company, error) {
    result := GetDbConn().Create(&company)
    if result.Error != nil {
        logrus.Errorf("Unable to save company, %s", result.Error)
        return nil, result.Error
    }

    var addedCompany models.Company
    result = GetDbConn().First(&addedCompany, "id = ?", company.ID)
    if result.Error != nil {
        logrus.Errorf("Unable to get company, %s", result.Error)
        return nil, result.Error
    }
    return &addedCompany, nil
}

func (s CompanyDbServiceImpl) GetCompanyById(companyId string) (*models.Company, error) {
    var company models.Company
    result := GetDbConn().First(&company, "id = ?", companyId)
    if result.Error != nil {
        logrus.Errorf("Unable to get company, %s", result.Error)
        return nil, result.Error
    }
    return &company, nil
}

func (s CompanyDbServiceImpl) GetAllCompanies() ([]*models.Company, error) {
    var companies []*models.Company
    result := GetDbConn().Preload("CompanyCertifications.Certification").Find(&companies)
    if result.Error != nil {
        logrus.Errorf("Unable to get companies, %s", result.Error)
        return nil, result.Error
    }
    return companies, nil
}


func (s CompanyDbServiceImpl) UpdateCompany() ([]*models.Company, error) {
    var companies []*models.Company
    result := GetDbConn().Find(&companies)
    if result.Error != nil {
        logrus.Errorf("Unable to get companies, %s", result.Error)
        return nil, result.Error
    }
    return companies, nil
}