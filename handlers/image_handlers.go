package handlers

import (
	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/model"
)

type ImageDbService struct{}

func (s ImageDbService) AddImage(request string, dbService database.ImageDbService) (*model.Image, error) {

	return &model.Image{}, nil
}
