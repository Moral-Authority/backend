package database

import (
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/models"
)

type UserDbService interface {
	AddNewUser(models.User) *models.User
	GetUser(uint) *models.User
	GetAllUsers() []models.User
	UpdateUser(uint, model.UpdateUser) *models.User
	AddUserFav(model.AddUserFav, models.Product) []models.Favourite
}
