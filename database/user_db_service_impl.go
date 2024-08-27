package database

import (
	"strconv"

	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
)

type UserDbServiceImpl struct{}

func (s UserDbServiceImpl) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    result := GetDbConn().Where("email = ?", email).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func (s UserDbServiceImpl) AddNewUser(newUser models.User) (*models.User, error) {
	result := GetDbConn().Create(&newUser)
	if result.Error != nil {
		logrus.Errorf("Unable to save user, %s", result.Error)
		return nil, result.Error
	}
	user, err := s.GetUser(strconv.Itoa(int(newUser.ID)))
	return user, err
}


func (s UserDbServiceImpl) GetUser(userId string) (*models.User, error) {
	var user models.User
	result := GetDbConn().First(&user, "id = ?", userId)
	if result.Error != nil {
		logrus.Errorf("Unable to get user, %s", result.Error)
		return nil, result.Error
	}
	var favs []models.Favorite
	results := GetDbConn().Model(&user).Association("Favorites").Find(&favs)
	if results != nil {
		logrus.Errorf("Unable to get categories, %s", results)
		return nil, results
	}
	user.Favorites = favs
	return &user, nil
}

func (s UserDbServiceImpl) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	result := GetDbConn().Find(&users)
	if result.Error != nil {
		logrus.Errorf("Unable to get all user, %s", result.Error)
		return nil, result.Error
	}
	return users, nil
}

func (s UserDbServiceImpl) UpdateUser(userId string, request model.UpdateUser) (*models.User, error) {
	user, err := s.GetUser(userId)
	if err != nil {
		logrus.Errorf("unable to get user")
		return nil, err
	}
	result := GetDbConn().Model(&user).Updates(models.User{
            Email: *request.Email,
			PasswordHash: *request.Password,
	})
	if result.Error != nil {
		logrus.Errorf("Unable to update user, %s", result.Error)
		return nil, result.Error
	}
	updatedUser, err := s.GetUser(userId)
	return updatedUser, err
}

func (s UserDbServiceImpl) AddUserFav(request model.AddUserFav, product models.Product) ([]models.Favorite, error) {
	id, err := StringToUint(request.UserID)
	if err != nil {
		logrus.Errorf("Unable to convert id, %s", request.UserID)
		return nil, err
	}
	user, err := s.GetUser(strconv.Itoa(int(id)))
	if err != nil {
		logrus.Errorf("unable to get user")
		return nil, err
	}
	err = GetDbConn().Model(&user).Association("Favorites").Append(models.Favorite{
		UserRefer: id,
		Product:   product,
	})
	if err != nil {
		logrus.Errorf("unable to append fav")
		return nil, err
	}
	updatedUser, err := s.GetUser(strconv.Itoa(int(id)))
	if err != nil {
		return nil, err
	}
	return updatedUser.Favorites, nil
}

func (s UserDbServiceImpl) RemoveUserFav(request model.RemoveUserFav, product models.Product) ([]models.Favorite, error) {
	id, err := StringToUint(request.UserID)
	if err != nil {
		logrus.Errorf("Unable to convert id, %s", request.UserID)
		return nil, err
	}
	user, err := s.GetUser(strconv.Itoa(int(id)))
	if err != nil {
		logrus.Errorf("unable to get user")
		return nil, err
	}

	// Find the favorite to be removed
	var favorite models.Favorite
	err = GetDbConn().Model(&user).Association("Favorites").Find(&favorite, "product_id = ?", product.ID)
	if err != nil {
		logrus.Errorf("unable to find fav")
		return nil, err
	}

	// Remove the favorite
	err = GetDbConn().Model(&user).Association("Favorites").Delete(&favorite)
	if err != nil {
		logrus.Errorf("unable to delete fav")
		return nil, err
	}

	updatedUser, err := s.GetUser(strconv.Itoa(int(id)))
	if err != nil {
		return nil, err
	}
	return updatedUser.Favorites, nil
}
