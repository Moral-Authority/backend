package handlers

import (
	"errors"
	"fmt"
	"os"

	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

// Secret key for signing the JWT token
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func (s UserService) AddNewUser(request model.NewUser, dbService database.UserDbService) (*model.User, error) {
	// 1: Hash the password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// 2: Create the user
	user := models.User{
		Email:        request.Email,
		Phone:        request.Phone,
		PasswordHash: string(hashedPassword),
		Favorites:    []models.Favorite{},
	}

	// 3: Save the user in the database
	savedUser, err := dbService.AddNewUser(user)
	if err != nil {
		return nil, err
	}

	return toUserResponse(*savedUser), nil
}

func (s UserService) Login(request model.LoginUser, dbService database.UserDbService) (string, *model.User, error) {
	// 1. Fetch the user from the database by email
	user, err := dbService.GetUserByEmail(request.Email)
	if err != nil {
		return "", nil, errors.New("user not found")
	}

	// 2. Compare the provided password with the stored password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return "", nil, errors.New("incorrect password")
	}

	// 3. Generate a JWT token
	tokenString, err := generateJWTToken(user)
	if err != nil {
		return "", nil, errors.New("failed to generate token")
	}

	// 4. Return the token and user details
	return tokenString, toUserResponse(*user), nil
}

func (s UserService) UpdateUser(request model.UpdateUser, dbService database.UserDbService) (*model.User, error) {
	updatedUser, err := dbService.UpdateUser(request.UserID, request)
	if err != nil {
		return nil, err
	}
	if updatedUser == nil {
		return nil, errors.New("unable to update user in db")
	}
	return toUserResponse(*updatedUser), nil
}

func (s UserService) GetUserById(userId string, dbService database.UserDbService) (*model.User, error) {
	user, err := dbService.GetUser(userId)
	if err != nil || user == nil {
		return nil, errors.New(fmt.Sprintf("unable to get user from db %s", err))
	}

	return toUserResponse(*user), nil
}

func (s UserService) GetUsers(dbService database.UserDbService) ([]*model.User, error) {
	users, err := dbService.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var response []*model.User
	for _, e := range users {
		user := toUserResponse(*e)
		response = append(response, user)
	}
	return response, nil
}

func (s UserService) ToggleUserFav(request model.ToggleUserFav, userDbService database.UserDbService, productDbService database.ProductDbService) ([]*model.Favorite, error) {

	// validate user
	user, err := userDbService.GetUser(request.UserID)
	if err != nil || user == nil {
		return nil, errors.New(fmt.Sprintf("unable to get user from db %s", err))
	}

	// validate product
	product, err := productDbService.GetProductByID(request.ProductID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to get product from db %s", err))
	}

	// check if favorite exists
	userFav, err := userDbService.GetUserFav(user.ID, product.ID)
	if err != nil {
		return nil, err
	}

	if userFav == nil {
		// add favorite
		_, err := userDbService.AddUserFav(request)
		if err != nil {
			return nil, err
		}

	} else {
		// remove favorite
		err := userDbService.RemoveUserFav(request)
		if err != nil {
			return nil, err
		}

	}

	favs, err := userDbService.GetAllUserFavs(user.ID)
	if err != nil {
		return nil, err
	}

	return toFavsResponse(favs), nil
}
