package database

import (
	"github.com/howstrongiam/backend/models"
	"github.com/sirupsen/logrus"
)

type ImageDbServiceImpl struct{}

func (s ImageDbServiceImpl) AddImage(image models.Image) *models.Image {
	result := GetDbConn().Create(&image)
	if result.Error != nil {
		logrus.Errorf("Unable to save image, %s", result.Error)
		return nil
	}
	var addedImage models.Image
	result = GetDbConn().First(&addedImage, "id = ?", image.ID)
	if result.Error != nil {
		logrus.Errorf("Unable to get image, %s", result.Error)
		return nil
	}
	return &addedImage
}
