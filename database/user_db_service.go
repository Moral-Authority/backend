package database

import (
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
)

type UserDbService interface {
	AddNewUser(models.User) *models.User
	GetUser(string) *models.User
	GetAllUsers() []models.User
	UpdateUser(string, model.UpdateUser) *models.User
	AddUserFav(model.AddUserFav, models.Product) []models.Favourite
}
