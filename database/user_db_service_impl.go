package database

import (
	"strconv"

	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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

func (s UserDbServiceImpl) AddUserFav(request model.ToggleUserFav) (*models.Favorite, error) {

	userID, err := StringToUint(request.UserID)
	if err != nil {
		logrus.Errorf("Unable to convert user id string to uint, %s", request.UserID)
		return nil, err
	}

	productID, err := StringToUint(request.ProductID)
	if err != nil {
		logrus.Errorf("Unable to convert product id string to uint, %s", request.UserID)
		return nil, err
	}


	var fav models.Favorite
	fav.UserRefer = userID
	fav.ProductId = productID
	
	result := GetDbConn().Create(&fav)
	if result.Error != nil {
		logrus.Errorf("Unable to save user, %s", result.Error)
		return nil, result.Error
	}

	return &fav, nil
}

func (s UserDbServiceImpl) RemoveUserFav(request model.ToggleUserFav) error {
	
	var favorite models.Favorite
	result := GetDbConn().Where("user_id = ? AND product_id = ?", request.UserID, request.ProductID).First(&favorite)
	if result.Error != nil {
		logrus.Errorf("Error finding favorite: %v", result.Error)
		return result.Error
	}

	result = GetDbConn().Delete(&favorite)
	if result.Error != nil {
		logrus.Errorf("Unable to delete favorite, %s", result.Error)
		return result.Error
	}
	return nil
}


func (s UserDbServiceImpl) GetUserFav(userID uint, productID uint) (*models.Favorite, error) {

	var favorite models.Favorite
	err := GetDbConn().Where("user_id = ? AND product_id = ?", userID, productID).First(&favorite).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		logrus.Errorf("Error finding favorite: %v", err)
		return nil, err
	}


	return &favorite, nil 
}

func (s UserDbServiceImpl) GetAllUserFavs(userID uint) ([]*models.Favorite, error) {
    var favs []*models.Favorite

    err := GetDbConn().Where("user_refer = ?", userID).
        Preload("Product").Find(&favs).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, nil
        }
        logrus.Errorf("Error finding favorites: %v", err)
        return nil, err
    }

    return favs, nil
}
