package handlers

import (
    "errors"
    "github.com/Moral-Authority/backend/database"
    "github.com/Moral-Authority/backend/graph/model"
    "github.com/Moral-Authority/backend/models" 
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
    savedUser, err := dbService.AddNewUser(user)
    if err != nil {
        return nil, err
    }
    // 3: return response
    if savedUser != nil {
        return toUserResponse(*savedUser), nil
    } else {
        return nil, errors.New("unable to save user in db")
    }
}

func (s UserService) UpdateUser(request model.UpdateUser, dbService database.UserDbService) (*model.User, error) {
    updatedUser,err := dbService.UpdateUser(request.UserID, request)
    if err != nil {
        return nil, err
    }
    if updatedUser == nil {
        return nil, errors.New("unable to update user in db")
    }
    return toUserResponse(*updatedUser), nil
}

func (s UserService) AddUserFav(request model.AddUserFav, userDbService database.UserDbService, productDbService database.ProductDbService) ([]*model.Favourite, error) {
    product, err := productDbService.GetProductByID(request.ProductID)
    if err != nil {
        return nil, err
    }

    if product == nil {
        return nil, errors.New("unable to get product")
    }

    addedFav,err := userDbService.AddUserFav(request, *product)
    if err != nil {
        return nil, err
    }
    
    return toFavsResponse(addedFav), nil
}

func (s UserService) RemoveUserFav(request model.AddUserFav, userDbService database.UserDbService, productDbService database.ProductDbService) ([]*model.Favourite, error) {
    product, err := productDbService.GetProductByID(request.ProductID)
    if err != nil {
        return nil, err
    }
    if product == nil {
        return nil, errors.New("unable to get product")
    }
    addedFav,err := userDbService.RemoveUserFav(request, *product)
    if err != nil {
        return nil, err
    }
    return toFavsResponse(addedFav), nil
}

func (s UserService) GetUserById(userId string, dbService database.UserDbService) (*model.User, error) {
    user,err := dbService.GetUser(userId)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, errors.New("unable to get user from db")
    }
    return toUserResponse(*user), nil
}

func (s UserService) GetUsers(dbService database.UserDbService) ([]*model.User, error) {
    users ,err:= dbService.GetAllUsers()
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