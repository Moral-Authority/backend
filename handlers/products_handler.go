package handlers

import (
	"errors"
	"fmt"

	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
)

type ProductService struct{}

func (s ProductService) AddNewProduct(request model.AddProductRequest, productDbService database.ProductDbService, imageDbService database.ImageDbService, certificationService database.CertificationDbService) (*model.Product, error) {

	// _, err := database.StringToUint(request.UserID)
	// if err != nil {
	// 	return nil, err
	// }

	product := models.Product{
		Url:         request.PurchaseInfo.Link,
		Description: request.Description,
		Title:       request.Title,
	}

	addedProduct, err := productDbService.AddProduct(product)
	if err != nil || addedProduct == nil {
		return nil, errors.New("unable to save product to db")
	}

	for _, i := range request.ImageLinks {

		image := models.Image{
			Url: *i,
		}

		addedImage, err := imageDbService.AddImage(image)
		if err != nil || addedImage == nil {
			return nil, errors.New("unable to save image to db")
		}
	}

	for _, c := range request.Certifications {

		foundCert, err := certificationService.GetCertificationById(*c.CertificationID)
		if err != nil  {
			return nil, fmt.Errorf("unable to find certification number %d", c.CertificationID)
		}

		cert := models.ProductCertification{
			ProductID:	   addedProduct.ID,
			CertificationID: foundCert.ID,
		}

		addedProductCert, err := productDbService.AddProductCertification(cert)
		if err != nil || addedProductCert == nil {
			return nil, errors.New("unable to save product certification to db")
		}
	}

	return toProductResponse(addedProduct), nil
}

func (s ProductService) UpdateProduct(request model.UpdateProductRequest, productDbService database.ProductDbService, imageDbService database.ImageDbService, certificationService database.CertificationDbService) (*model.Product, error) {
	dbService := database.ProductDbServiceImpl{}

	product, err := dbService.UpdateProduct(request)
	if err != nil {
		return nil, err
	}

	return toProductResponse(product), nil
}

func (s ProductService) GetProductByID(productId string, productDbService database.ProductDbService) (*model.Product, error) {
	dbService := database.ProductDbServiceImpl{}

	product, err := dbService.GetProductByID(productId)
	if err != nil {
		return nil, err
	}

	return toProductResponse(product), nil
}

func (s ProductService) GetAllProducts(productDbService database.ProductDbService) ([]*model.Product, error) {
	dbService := database.ProductDbServiceImpl{}
	var result []*model.Product

	products, err := dbService.GetAllProducts()
	if err != nil {
		return nil, err
	}

	for _, p := range products {
		result = append(result, toProductResponse(p))
	}

	return result, nil
}

func (s ProductService) GetProductsByFilter(filters map[string]interface{}, productDbService database.ProductDbService) ([]*model.Product, error) {
	dbService := database.ProductDbServiceImpl{}
	var result []*model.Product

	products, err := dbService.GetAllProducts()
	if err != nil {
		return nil, err
	}

	for _, p := range products {
		result = append(result, toProductResponse(p))
	}

	return result, nil
}
