package handlers

import (
	"errors"
	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/models"
)

type ProductService struct{}

func (s ProductService) AddNewProduct(
	request model.AddProductRequest,
	productDbService database.ProductDbService,
	imageDbService database.ImageDbService,
	certDbService database.CertificatesDbService,
) (*model.Product, error) {
	styleId, err := database.StringToUint("request.StyleID")
	if err != nil {
		return nil, err
	}
	userId, err := database.StringToUint("request.UserID")
	if err != nil {
		return nil, err
	}
	image := models.Image{
		ImageLocation: "request.ImageLocation",
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
	product := models.Product{
		Url:             "request.URL",
		Description:     request.Description,
		Title:           request.Title,
		UserId:          userId,
		ImageId:         addedImage.ID,
		CertificationId: addedCert.ID,
		StyleRefer:      styleId,
	}
	addedProduct := productDbService.AddProduct(product)
	if addedProduct == nil {
		return nil, errors.New("unable to save product to db")
	}
	return toProductResponse(*addedProduct), nil
}
