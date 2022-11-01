package repos

import (
	"github.com/howstrongiam/backend/database"
	"github.com/howstrongiam/backend/models"
)

func InsertUser(user models.User) {
	database := database.GetDatabase()

	//Insert into user
	database.Create(user)
}
