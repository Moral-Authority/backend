package database

import (
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
)

type UserDbService interface {
	AddNewUser(models.User) (*models.User, error)
	GetUser(string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(string, model.UpdateUser) (*models.User, error)
	AddUserFav(model.ToggleUserFav) (*models.Favorite, error)
	RemoveUserFav(model.ToggleUserFav) error
	GetUserFav(uint, uint) (*models.Favorite, error)
	GetUserByEmail(email string) (*models.User, error)
}
