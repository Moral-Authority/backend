package database

import (
    "github.com/Moral-Authority/backend/graph/model"
    "github.com/Moral-Authority/backend/models"
)

type UserDbService interface {
    AddNewUser(models.User) (*models.User, error)
    GetUser(string) (*models.User,error)
    GetAllUsers() ([]*models.User,error) // Corrected return type
    UpdateUser(string, model.UpdateUser)( *models.User,error)
    AddUserFav(model.AddUserFav, models.Product) ([]models.Favourite,error)
}