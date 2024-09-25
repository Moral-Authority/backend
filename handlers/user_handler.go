package handlers

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/Moral-Authority/backend/database"
	"github.com/Moral-Authority/backend/graph/model"
	"github.com/Moral-Authority/backend/models"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

// Secret key for signing the JWT token
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func (s UserService) AddNewUserHandler(request model.NewUser, dbService database.UserDbService) (*model.User, error) {
	// 1: Hash the password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// 2: Generate email verification token
	verificationToken, err := generateVerificationToken(request.Email)
	if err != nil {
		return nil, errors.New("failed to generate verification token")
	}

	// 3: Create the user with Verified set to false
	user := models.User{
		Email:             request.Email,
		Phone:             request.Phone,
		PasswordHash:      string(hashedPassword),
		Favorites:         []models.Favorite{},
		Verified:          false,             // Initially unverified
		VerificationToken: verificationToken, // Store the verification token
	}

	// 5: Send verification email
	err = sendVerificationEmail(user.Email, verificationToken)
	if err != nil {
		return nil, errors.New("failed to send verification email")
	}

	// 3: Save the user in the database
	savedUser, err := dbService.AddNewUser(user)
	if err != nil {
		return nil, err
	}

	return toUserResponse(*savedUser), nil
}

func (s UserService) LoginHandler(request model.LoginUser, dbService database.UserDbService) (string, *model.User, error) {
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

	if !user.Verified {
		return "", nil, errors.New("email not verified")
	}

	// 3. Generate a JWT token
	tokenString, err := generateJWTToken(user)
	if err != nil {
		return "", nil, errors.New("failed to generate token")
	}

	// 4. Return the token and user details
	return tokenString, toUserResponse(*user), nil
}

func (s UserService) LogoutHandler(userID string, userDbService database.UserDbService) (*model.User, error) {

	// validate user
	user, err := userDbService.GetUser(userID)
	if err != nil || user == nil {
		return nil, fmt.Errorf("unable to get user from db %s", err)
	}

	return nil, nil
}

func (s UserService) UpdateUserHandler(request model.UpdateUser, dbService database.UserDbService) (*model.User, error) {
	updatedUser, err := dbService.UpdateUser(request.UserID, request)
	if err != nil {
		return nil, err
	}
	if updatedUser == nil {
		return nil, errors.New("unable to update user in db")
	}
	return toUserResponse(*updatedUser), nil
}

func (s UserService) GetUserByIdHandler(userId string, dbService database.UserDbService) (*model.User, error) {
	user, err := dbService.GetUser(userId)
	if err != nil || user == nil {
		return nil, fmt.Errorf("unable to get user from db %s", err)
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

func (s UserService) ToggleUserFavHandler(request model.ToggleUserFav, userDbService database.UserDbService, productDbService database.ProductDbService) ([]*model.Favorite, error) {

	// Validate user
	user, err := userDbService.GetUser(request.UserID)
	if err != nil || user == nil {
		return nil, fmt.Errorf("unable to get user from db: %v", err)
	}

	pd, isDepartment := IsStringValidProductDepartment(request.ProductDepartment)
	if !isDepartment {
		return nil, fmt.Errorf("invalid department type: %s", request.ProductDepartment)
	}

	// Validate product
	product, err := ProductService{}.GetProductByIDHandler(request.ProductID, pd.ToInt(), productDbService)
	if err != nil {
		return nil, fmt.Errorf("unable to get product from db: %v", err)
	}

	// Convert product ID to uint
	productIdUint, err := database.StringToUint(product.ID)
	if err != nil {
		return nil, err
	}

	// Check if favorite exists
	userFav, err := userDbService.GetUserFav(user.ID, productIdUint, pd.ToInt())
	if err != nil {
		return nil, err
	}

	if userFav == nil {
		// Add favorite
		logrus.Infof("adding fav: %d", pd.ToInt())
		_, err := userDbService.AddUserFav(request, pd.ToInt())
		if err != nil {
			return nil, err
		}
	} else {
		// Remove favorite
		logrus.Infof("removing fav: %d", pd.ToInt())
		err := userDbService.RemoveUserFav(request, pd.ToInt())
		if err != nil {
			return nil, err
		}
	}

	// Return updated list of favorites
	return s.GetAllUserFavsHandler(request.UserID, userDbService, productDbService)
}

func (s UserService) GetAllUserFavsHandler(userID string, userDbService database.UserDbService, productDbService database.ProductDbService) ([]*model.Favorite, error) {
	userId, err := database.StringToUint(userID)
	if err != nil {
		return nil, err
	}

	// Fetch all user favorites from the favorites table
	favs, err := userDbService.GetAllUserFavs(userId)
	if err != nil {
		return nil, err
	}

	// Channel to collect the results
	resultChan := make(chan *model.Favorite, len(favs))
	var wg sync.WaitGroup

	// Use Goroutines to fetch the details for each favorite concurrently
	for _, fav := range favs {
		wg.Add(1)
		go func(fav *models.Favorite) {
			defer wg.Done()

			var product interface{}
			var err error

			// Determine the correct table/service based on ProductDepartment
			switch fav.ProductDepartment {
			case HomeGardenProductDepartment.ToInt():
				product, err = productDbService.GetHomeGardenProductByID(fav.ProductID)

			case ClothingAccessoriesProductDepartment.ToInt():
				product, err = productDbService.GetClothingAccessoriesProductByID(fav.ProductID)

			case HealthBathBeautyProductDepartment.ToInt():
				product, err = productDbService.GetHealthBathBeautyProductByID(fav.ProductID)

			case ToysKidsBabiesProductDepartment.ToInt():
				product, err = productDbService.GetToysKidsBabiesProductByID(fav.ProductID)

			default:
				logrus.Errorf("Unknown ProductDepartment %d for favorite ID %d", fav.ProductDepartment, fav.ID)
				return
			}

			if err != nil {
				logrus.Errorf("Error fetching product details for favorite ID %d: %v", fav.ID, err)
				return
			}

			// Create the Favorite response model
			resultChan <- toFavResponse(fav, product, ProductDepartment(fav.ProductDepartment))

		}(fav)
	}

	// Close the channel once all Goroutines are done
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results from the channel
	var results []*model.Favorite
	for result := range resultChan {
		results = append(results, result)
	}

	return results, nil
}

// RequestPasswordResetHandler handles the password reset request
func (s UserService) RequestPasswordResetHandler(email string, dbService database.UserDbService) error {
	// Check if the user exists
	user, err := dbService.GetUserByEmail(email)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	// Generate a reset token
	token, err := generateResetToken(user.Email)
	if err != nil {
		return errors.New("failed to generate reset token")
	}

	// Send the reset link to the user's email
	resetLink := "https://moralauthority.co/reset-password?token=" + token
	err = sendResetEmail(user.Email, resetLink)
	if err != nil {
		return errors.New("failed to send reset email")
	}

	return nil
}

// VerifyEmailHandler verifies the token and updates the user's status
func (s UserService) VerifyEmailHandler(token string, dbService database.UserDbService) (bool, error) {

	claims := &jwt.StandardClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !parsedToken.Valid {
		return false, errors.New("invalid or expired token")
	}


	user, err := dbService.GetUserByEmail(claims.Subject)
	if err != nil || user == nil {
		return false, fmt.Errorf("user not found: %v", err)
	}

	user.Verified = true
	err = dbService.UpdateUserVerification(user) // Assuming you have a method to update verification status
	if err != nil {
		return false, errors.New("failed to update user verification status")
	}

	return true, nil
}
