package handlers

import (
	"errors"

	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models" // Assuming this is the correct import path for your models
)

type ProductService struct{}

func (s ProductService) AddNewProduct(request model.AddProductRequest, productDbService database.ProductDbService, imageDbService database.ImageDbService, certificationService database.CertificationDbService) (*model.Product, error) {

	// _, err := database.StringToUint(request.UserID)
	// if err != nil {
	// 	return nil, err
	// }


	product := models.Product{
		Url:         request.PurchaseInfo.Link, // Assuming Url is in PurchaseInfo
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

	// TODO for each image to product_images table

	// TODO for each certificate id search db for cert and add to product_certs relational table

	return toProductResponse(*addedProduct), nil
}


func (s ProductService) UpdateProduct(request model.AddProductRequest, productDbService database.ProductDbService, imageDbService database.ImageDbService, certificationService database.CertificationDbService) (*model.Product, error) {
	dbService := database.ProductDbServiceImpl{}

	product, err := dbService.UpdateProduct(filters)
	if err != nil {
		return nil, err
	}

	return toProductResponse(product), nil
}


func (s ProductService) GetProductByID(request model.AddProductRequest, productDbService database.ProductDbService) (*model.Product, error) {
	dbService := database.ProductDbServiceImpl{}

	product, err := dbService.GetProductByID(filters)
	if err != nil {
		return nil, err
	}

	return toProductResponse(product), nil
}


func (s ProductService) GetAllProducts(request model.AddProductRequest, productDbService database.ProductDbService) ([]*model.Product, error) {
	dbService := database.ProductDbServiceImpl{}
	var result []*model.Product

	products, err := dbService.GetAllProducts(filters)
	if err != nil {
		return nil, err
	}

	for _, p := range products {
		result = append(result,  toProductResponse(p))
	}

	return result, nil
}
