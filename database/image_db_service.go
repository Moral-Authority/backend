package database

import "github.com/Moral-Authority/backend/models"

type ImageDbService interface {
	AddImage(image models.Image) (*models.Image, error)
	UpdateImage(image models.Image) (*models.Image, error)
	GetImageById(imageId string) (*models.Image, error)
}
