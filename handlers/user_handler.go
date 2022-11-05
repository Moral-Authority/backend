package handlers

import (
	"errors"
	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/models"
)

type UserService struct{}

func (s UserService) AddNewUser(request model.NewUser, dbService database.UserDbService) (*model.User, error) {
	// 1: save login credentials
	salt := generateRandomSalt(16)
	credentials := models.LoginCredentials{
		Email:        request.Email,
		PasswordHash: hashPassword(request.Password, salt),
		Salt:         salt,
	}
	// 2: save user
	user := models.User{
		FirstName:        request.FirstName,
		LastName:         request.LastName,
		LoginCredentials: credentials,
		Favourites:       []models.Favourite{},
	}
	savedUser := dbService.AddNewUser(user)
	// 3: return response
	if savedUser != nil {
		return toUserResponse(*savedUser), nil
	} else {
		return nil, errors.New("unable to save user in db")
	}
}

func (s UserService) UpdateUser(request model.UpdateUser, dbService database.UserDbService) (*model.User, error) {
	updatedUser := dbService.UpdateUser(request.UserID, request)
	if updatedUser == nil {
		return nil, errors.New("unable to update user in db")
	}
	return toUserResponse(*updatedUser), nil
}

func (s UserService) AddUserFav(request model.AddUserFav, userDbService database.UserDbService, productDbService database.ProductDbService) ([]*model.Favourite, error) {
	product := productDbService.GetProduct(request.ProductID)
	if product == nil {
		return nil, errors.New("unable to get product")
	}
	addedFav := userDbService.AddUserFav(request, *product)
	return toFavsResponse(addedFav), nil
}

func (s UserService) GetUserById(userId string, dbService database.UserDbService) (*model.User, error) {
	user := dbService.GetUser(userId)
	if user == nil {
		return nil, errors.New("unable to get user from db")
	}
	return toUserResponse(*user), nil
}

func (s UserService) GetUsers(dbService database.UserDbService) ([]*model.User, error) {
	users := dbService.GetAllUsers()
	var response []*model.User
	for _, e := range users {
		user := toUserResponse(e)
		response = append(response, user)
	}
	return response, nil
}
