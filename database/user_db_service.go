package database

import (
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/models"
)

type UserDbService interface {
	AddNewUser(models.User) *models.User
	GetUser(string) *models.User
	GetAllUsers() []models.User
	UpdateUser(string, model.UpdateUser) *models.User
	AddUserFav(model.AddUserFav, models.Product) []models.Favourite
}
