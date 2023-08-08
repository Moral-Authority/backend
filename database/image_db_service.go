package database

import "github.com/howstrongiam/backend/models"

type ImageDbService interface {
	AddImage(image models.Image) *models.Image
	//AddImages(image models.Image) *models.Image
	//UpdateImage(image models.Image)*models.Image
}
