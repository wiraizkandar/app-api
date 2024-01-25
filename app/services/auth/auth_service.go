package authservice

import (
	"errors"

	"github.com/officemaid/app-api/app/models"
	"github.com/officemaid/app-api/database"
	"golang.org/x/crypto/bcrypt"
)

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type AuthenticatedUserInfo struct {
	User        models.User `json:"user"`
	AccessToken AccessToken `json:"token"`
}

const HASH_LENGTH int = 12

func AuthenticateUser(username string, password string) (AuthenticatedUserInfo, error) {

	db := database.Init()

	var userData models.User

	userResult := db.Where("email = ?", username).First(&userData)

	if userResult.Error != nil {
		return AuthenticatedUserInfo{
			User:        userData,
			AccessToken: AccessToken{},
		}, errors.New("no user found")
	}

	// check is entered password is match with db password
	match := isPasswordMatch(userData.Password, password)

	if !match {
		return AuthenticatedUserInfo{}, errors.New("password does not match")
	}

	// register in authtoken table
	accessToken, _ := createAccessToken(userData)

	return AuthenticatedUserInfo{
		User:        userData,
		AccessToken: accessToken,
	}, nil

}

/**
 * Check if password match
 * @param userPassword string
 * @param password string
 * @return bool
 */
func isPasswordMatch(userPassword string, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))

	return err == nil
}

func createAccessToken(user models.User) (AccessToken, error) {

	return AccessToken{
		AccessToken:  user.Id,
		RefreshToken: user.Id,
		ExpiresIn:    3600,
	}, nil
}
