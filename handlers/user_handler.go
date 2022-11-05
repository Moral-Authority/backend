package handlers

import (
	"github.com/howstrongiam/backend/graph/model"
	//userModels "github.com/howstrongiam/backend/models"
)

type UserService struct{}

func (s UserService) AddNewUser(request model.NewUser) (*model.User, error) {
	panic("NOT IMPL")
}

func (s UserService) UpdateUser(request model.UpdateUser) (*model.User, error) {
	panic("NOT IMPL")
}

func (s UserService) AddUserFav(request model.AddUserFav) (*model.Favourite, error) {
	panic("NOT IMPL")
}

func (s UserService) GetUserById(userId string) (*model.User, error) {
	panic("NOT IMPL")
}

func (s UserService) GetUsers() ([]*model.User, error) {
	panic("NOT IMPL")
}
