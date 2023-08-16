package handlers

import (
	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
)

type ImageDbService struct{}

func (s ImageDbService) AddImage(request string, dbService database.ImageDbService) (*model.Image, error) {

	return &model.Image{}, nil
}
