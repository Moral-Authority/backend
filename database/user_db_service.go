package database

import "github.com/howstrongiam/backend/models"

type UserDbService interface {
	AddNewUser(models.User) *models.User
	GetUser(string) *models.User
	GetAllUsers() []models.User
	DeleteUser(string) bool
	//UpdateUser(string, models.UpdateUserRequest) *models.User
}
