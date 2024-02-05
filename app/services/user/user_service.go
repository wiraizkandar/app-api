package user

import (
	"errors"

	"github.com/officemaid/app-api/app/models"
	"github.com/officemaid/app-api/database/mysql"
)

func GetUser(userId string) (models.User, error) {

	db := mysql.Init()

	var userData models.User

	userResult := db.Where("id = ?", userId).First(&userData)

	if userResult.Error != nil {
		return models.User{}, errors.New("no record found")
	}

	return userData, nil

}
