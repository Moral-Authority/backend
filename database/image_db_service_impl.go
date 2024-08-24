package database

import (
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
)

type ImageDbServiceImpl struct{}

func (s ImageDbServiceImpl) AddImage(image models.Image) (*models.Image, error) {
	result := GetDbConn().Create(&image)
	if result.Error != nil {
		logrus.Errorf("Unable to save image, %s", result.Error)
		return nil, result.Error
	}
	var addedImage models.Image
	result = GetDbConn().First(&addedImage, "id = ?", image.ID)
	if result.Error != nil {
		logrus.Errorf("Unable to get image, %s", result.Error)
		return nil, result.Error
	}
	return &addedImage, nil 
}


func (s ImageDbServiceImpl) UpdateImage(image models.Image) (*models.Image, error) {
	var currentImage models.Image

    result := GetDbConn().First(&currentImage, "id = ?", image.ID)
    if result.Error != nil {
        logrus.Errorf("Unable to find image, %s", result.Error)
        return nil, result.Error
    }

    // Update the existing certification with the new data
    result = GetDbConn().Model(&currentImage).Updates(image)
    if result.Error != nil {
        logrus.Errorf("Unable to update certification, %s", result.Error)
        return nil, result.Error
    }

    return &currentImage, nil
}

func (s ImageDbServiceImpl) GetImageById(imageId string) (*models.Image, error) {
    var image models.Image
    result := GetDbConn().First(&image, "id = ?", imageId)
    if result.Error != nil {
        logrus.Errorf("Unable to get image, %s", result.Error)
        return nil, result.Error
    }
    return &image, nil
}