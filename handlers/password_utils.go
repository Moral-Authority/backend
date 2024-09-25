package handlers

import (
	"time"

	"github.com/Moral-Authority/backend/models"
	jwt "github.com/golang-jwt/jwt/v4"
)

// Function to generate a JWT token
func generateJWTToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
