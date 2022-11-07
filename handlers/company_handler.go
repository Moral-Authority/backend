package handlers

import (
	"errors"
	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/models"
)

type CompanyService struct{}

func (s CompanyService) AddCompany(request model.AddCompanyRequest, dbService database.CompanyDbService,
	imageDbService database.ImageDbService,
	certDbService database.CertificatesDbService,
) (*model.Company, error) {
	userId, err := database.StringToUint(request.UserID)

	if err != nil {
		return nil, err
	}

	image := models.Image{
		ImageLocation: request.ImageLocation,
	}

	addedImage := imageDbService.AddImage(image)
	if addedImage == nil {
		return nil, errors.New("unable to save image to db")
	}

	cert := models.Certification{
		CertifyingCompany: "request.Certification.CertifyingCompany",
		CertName:          "request.Certification.CertName",
	}
	addedCert := certDbService.AddNewCertificate(cert)
	if addedCert == nil {
		return nil, errors.New("unable to save certificate to db")
	}
	company := models.Company{
		Url:             request.URL,
		Description:     request.Description,
		UserId:          userId,
		IsVerified:      request.IsVerified,
		ImageId:         addedImage.ID,
		CertificationId: addedCert.ID,
	}
	savedCompany := dbService.AddCompany(company)
	if savedCompany == nil {
		return nil, errors.New("unable to save company in db")
	}
	return toCompanyResponse(*savedCompany), nil
}

func (s CompanyService) GetCompanyById(companyId string, dbService database.CompanyDbService) (*model.Company, error) {
	company := dbService.GetCompanyById(companyId)
	if company == nil {
		return nil, errors.New("unable to get company from db")
	}
	return toCompanyResponse(*company), nil
}
