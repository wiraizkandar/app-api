package authservice

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/officemaid/app-api/app/models"
	"github.com/officemaid/app-api/database/cassandra"
	"github.com/officemaid/app-api/database/mysql"
	"golang.org/x/crypto/bcrypt"
)

type AccessToken struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    time.Time `json:"expires_in"`
}

type AuthenticatedUserInfo struct {
	User        models.User `json:"user"`
	AccessToken AccessToken `json:"token"`
}

type AccessTokenClaims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

/**
 * Authenticate user
 * @param username string
 * @param password string
 * @return AuthenticatedUserInfo
 */
func AuthenticateUser(username string, password string) (AuthenticatedUserInfo, error) {

	db := mysql.Init()

	var userData models.User

	userResult := db.Where("email = ?", username).First(&userData)

	if userResult.Error != nil {
		return AuthenticatedUserInfo{}, errors.New("no valid user found")
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

/**
 * Create access token
 * @param user models.User
 * @return AccessToken
 */
func createAccessToken(user models.User) (AccessToken, error) {

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &AccessTokenClaims{
		UserId:   user.Id,
		Username: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	accessToken, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))

	refreshToken := uuid.New().String()

	if err != nil {
		return AccessToken{}, errors.New("failed to create token")
	}

	// register token in cassandra database
	registerRefreshToken(user.Id, refreshToken, expirationTime)

	return AccessToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expirationTime,
	}, nil
}

func registerRefreshToken(userId string, refreshToken string, expirationTime time.Time) (bool, error) {

	db := cassandra.Init()

	defer db.Close()

	// Execute the SQL query
	if err := db.Query(`INSERT INTO access_tokens (user_id, access_token, expiry,is_revoke) VALUES (?, ?, ?, ?)`,
		userId,
		refreshToken,
		expirationTime,
		false,
	).Exec(); err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}
