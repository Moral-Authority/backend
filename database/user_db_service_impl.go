package database

import (
	"github.com/howstrongiam/backend/graph/model"
	"github.com/howstrongiam/backend/models"
	"github.com/sirupsen/logrus"
	"strconv"
)

type UserDbServiceImpl struct{}

func (s UserDbServiceImpl) AddNewUser(newUser models.User) *models.User {
	result := GetDbConn().Create(&newUser)
	if result.Error != nil {
		logrus.Errorf("Unable to save user, %s", result.Error)
		return nil
	}
	return s.GetUser(strconv.Itoa(int(newUser.ID)))
}

func (s UserDbServiceImpl) GetUser(userId string) *models.User {
	var user models.User
	result := GetDbConn().First(&user, "id = ?", userId)
	if result.Error != nil {
		logrus.Errorf("Unable to get user, %s", result.Error)
		return nil
	}
	var favs []models.Favourite
	results := GetDbConn().Model(&user).Association("Favourites").Find(&favs)
	if results != nil {
		logrus.Errorf("Unable to get categories, %s", results)
		return nil
	}
	user.Favourites = favs
	return &user
}

func (s UserDbServiceImpl) GetAllUsers() []models.User {
	var users []models.User
	result := GetDbConn().Find(&users)
	if result.Error != nil {
		logrus.Errorf("Unable to get all user, %s", result.Error)
		return nil
	}
	return users
}

func (s UserDbServiceImpl) UpdateUser(userId string, request model.UpdateUser) *models.User {
	// 1: get the user first
	user := s.GetUser(userId)
	if user == nil {
		logrus.Errorf("unable to get user")
		return nil
	}
	// 2: update user
	result := GetDbConn().Model(&user).Updates(models.User{
		FirstName: *request.FirstName,
		LastName:  *request.LastName,
	})
	if result.Error != nil {
		logrus.Errorf("Unable to update user, %s", result.Error)
		return nil
	}
	// 3: return updated user
	return s.GetUser(userId)
}

func (s UserDbServiceImpl) AddUserFav(request model.AddUserFav, product models.Product) []models.Favourite {
	// 1: get the user first
	id, err := StringToUint(request.UserID)
	if err == nil {
		logrus.Errorf("Unable to convert id, %s", request.UserID)
		return nil
	}
	user := s.GetUser(strconv.Itoa(int(id)))
	if user == nil {
		logrus.Errorf("unable to get user")
		return nil
	}
	err = GetDbConn().Model(&user).Association("Favourites").Append(models.Favourite{
		UserRefer: id,
		Product:   product,
	})
	if err != nil {
		logrus.Errorf("unable to append fav")
		return nil
	}
	updatedUser := s.GetUser(strconv.Itoa(int(id)))
	return updatedUser.Favourites
}
