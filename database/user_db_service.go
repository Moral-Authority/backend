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
	AddUserFav(model.ToggleUserFav, int) (*models.Favorite, error)
	RemoveUserFav(model.ToggleUserFav, int) error
	GetUserFav(uint, uint, int) (*models.Favorite, error)
	GetAllUserFavs(uint) ([]*models.Favorite, error)
	GetUserByEmail(email string) (*models.User, error)
}
