package handlers

import (
	"errors"

	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
)



type ImageDbService struct{}

func (s ImageDbService) AddImage(dbService database.ImageDbService, request model.AddImage) (*model.Image, error) {

	    // _, err := database.StringToUint(request.UserID)
    // if err != nil {
    //     return nil, err
    // }

	image := models.Image{
		Url: request.URL,
	}

	if  request.ProductID != nil {
		// validate product id 
	}

    savedImage, err := dbService.AddImage(image)
    if err != nil || savedImage == nil {
        return nil, errors.New("unable to save company in db")
    }
	return toImageResponse(*savedImage), nil
}

func (s ImageDbService) UpdateImage(dbService database.ImageDbService, request model.UpdateImage) (*model.Image, error) {
    image, err := dbService.GetImageById(request.ID)
    if err != nil || image == nil {
        return nil, errors.New("unable to get image from db")
    }
    return toImageResponse(*image), nil
}