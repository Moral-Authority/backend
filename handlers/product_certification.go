package handlers

import (
	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
)

func (s ProductService) AddProductCertification(dbService database.ProductDbService, input model.ProductCertificationInput) (*model.ProductCertification, error) {

	// _, err := database.StringToUint(request.UserID)
	// if err != nil {
	//     return nil, err
	// }

	// TODO validate product id
	// TODO validate product id

	// savedImage, err := dbService.AddProductCertification(request)
	// if err != nil || savedImage == nil {
	//     return nil, errors.New("unable to save product certification in db")
	// }
	return nil, nil
}

// func (s ProductService) DeleteProductCertification(dbService database.ImageDbService, request model.UpdateImage) (*model.Image, error) {
//     image, err := dbService.GetProductCertificationById(request.ID)
//     if err != nil || image == nil {
//         return nil, errors.New("unable to get image from db")
//     }
//     return toImageResponse(*image), nil
// }
