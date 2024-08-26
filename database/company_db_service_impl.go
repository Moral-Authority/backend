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


func (s CompanyDbServiceImpl) UpdateCompany(input models.Company) (*models.Company, error) {
    var currentCompany *models.Company
    result := GetDbConn().First(&currentCompany, "id = ?", input.ID)
    if result.Error != nil {
        logrus.Errorf("Unable to find image, %s", result.Error)
        return nil, result.Error
    }

    // Update the existing certification with the new data
    result = GetDbConn().Model(&currentCompany).Updates(input)
    if result.Error != nil {
        logrus.Errorf("Unable to update certification, %s", result.Error)
        return nil, result.Error
    }

    return currentCompany, nil
}

func (s CompanyDbServiceImpl) GetCompaniesByFilter(filters map[string]interface{}) ([]models.Company, error) {
    var companies []models.Company
    db := GetDbConn()

    query := ApplyFilters(db, filters)

    if err := query.Find(&companies).Error; err != nil {
        logrus.Errorf("Unable to get companies by filter, %s", err)
        return nil, err
    }

    return companies, nil
}

func (s CompanyDbServiceImpl) AddCompanyCertification(companyCertification models.CompanyCertification) (*models.CompanyCertification, error) {
    result := GetDbConn().Create(&companyCertification)
    if result.Error != nil {
        logrus.Errorf("Unable to save company, %s", result.Error)
        return nil, result.Error
    }

    var addedCompanyCertification models.CompanyCertification
    result = GetDbConn().First(&addedCompanyCertification, "id = ?", addedCompanyCertification.ID)
    if result.Error != nil {
        logrus.Errorf("Unable to get company, %s", result.Error)
        return nil, result.Error
    }
    return &addedCompanyCertification, nil
}
