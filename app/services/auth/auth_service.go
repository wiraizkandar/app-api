package authservice

import (
	"fmt"

	"github.com/officemaid/app-api/app/models"
	"github.com/officemaid/app-api/database"
)

func AuthenticateUser(username string, password string) models.User {

	db := database.Init()

	var userData models.User

	result := db.Where("email = ?", username).First(&userData)

	fmt.Println("Username: ", result.RowsAffected)

	return userData
}
