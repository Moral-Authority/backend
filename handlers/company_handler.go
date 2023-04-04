package handlers

import (
	"errors"
	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/model"
)

type CompanyService struct{}

func (s CompanyService) AddCompany(request model.AddCompany, dbService database.CompanyDbService,
	imageDbService database.ImageDbService,
	certDbService database.CertificationDbService,
) (*model.Company, error) {
	//userId, err := database.StringToUint(request.UserID)
	//if err != nil {
	//	return nil, err
	//}

	//if request.Logo.Valid {
	//
	//}
	//image := models.Image{
	//	ImageLocation: request.Logo,
	//}
	//
	//addedImage := imageDbService.AddImage(image)
	//if addedImage == nil {
	//	return nil, errors.New("unable to save image to db")
	//}
	//
	//cert := models.Certification{
	//	CertifyingCompany: "request.Certification.CertifyingCompany",
	//	CertName:          "request.Certification.CertName",
	//}
	//addedCert := certDbService.AddNewCertificate(cert)
	//if addedCert == nil {
	//	return nil, errors.New("unable to save certificate to db")
	//}
	//company := models.Company{
	//	Url:             request.URL,
	//	Description:     request.Description,
	//	UserId:          userId,
	//	IsVerified:      request.IsVerified,
	//	ImageId:         addedImage.ID,
	//	CertificationId: addedCert.ID,
	//}
	//savedCompany := dbService.AddCompany(company)
	//if savedCompany == nil {
	//	return nil, errors.New("unable to save company in db")
	//}
	//return toCompanyResponse(*savedCompany), nil
	return nil, nil
}

func (s CompanyService) GetCompanyById(companyId string, dbService database.CompanyDbService) (*model.Company, error) {
	company := dbService.GetCompanyById(companyId)
	if company == nil {
		return nil, errors.New("unable to get company from db")
	}
	return toCompanyResponse(*company), nil
}
