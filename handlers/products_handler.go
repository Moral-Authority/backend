package handlers

import (
	"errors"

	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models" // Assuming this is the correct import path for your models
)

type ProductService struct{}

func (s ProductService) AddNewProduct(request model.AddProductRequest, productDbService database.ProductDbService, imageDbService database.ImageDbService,) (*model.Product, error) {

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
			ImageLocation: *i, // Assuming ImageLocation is a string
		}

		addedImage := imageDbService.AddImage(image)
		if addedImage == nil {
			return nil, errors.New("unable to save image to db")
		}
	}

	// TODO for each image to product_images table

	// TODO for each certificate id search db for cert and add to product_certs relational table

	return toProductResponse(*addedProduct), nil
}
