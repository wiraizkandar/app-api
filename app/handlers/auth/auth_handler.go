package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/officemaid/app-api/app/data/response"
	authservice "github.com/officemaid/app-api/app/services/auth"
)

type AuthenticateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
	UserId       string `json:"user_id"`
}

/**
 * @api {post} /api/v1/auth/authenticate Sign In
 */
func Authenticate(c *fiber.Ctx) error {

	request := new(AuthenticateRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	userData, err := authservice.AuthenticateUser(request.Username, request.Password)

	if err != nil {
		// return
		return c.Status(fiber.StatusUnauthorized).JSON(response.ApiResponse{
			Status:  response.ErrorStatus,
			Message: "Authentication failed",
		})
	}

	// set access token cookie
	cookieAccessToken := fiber.Cookie{
		Name:     "jwt-access-token",
		Value:    userData.AccessToken.AccessToken,
		Expires:  userData.AccessToken.ExpiresIn,
		HTTPOnly: true,
	}

	c.Cookie(&cookieAccessToken)

	// set refresh token cookie
	cookieRefreshToken := fiber.Cookie{
		Name:     "jwt-refresh-token",
		Value:    userData.AccessToken.RefreshToken,
		Expires:  userData.AccessToken.ExpiresIn,
		HTTPOnly: true,
	}

	c.Cookie(&cookieRefreshToken)

	// return success response
	return c.Status(fiber.StatusOK).JSON(response.ApiResponse{
		Status:  response.SuccessStatus,
		Message: "User authenticated successfully",
		Data:    userData.User,
	})

}

func RefreshToken(c *fiber.Ctx) error {

	request := new(RefreshTokenRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	userData, err := authservice.RefreshToken(request.UserId, request.RefreshToken)

	if err != nil {
		// return
		return c.Status(fiber.StatusUnauthorized).JSON(response.ApiResponse{
			Status:  response.ErrorStatus,
			Message: err.Error(),
		})
	}

	// set access token cookie
	cookieAccessToken := fiber.Cookie{
		Name:     "access-token",
		Value:    userData.AccessToken.AccessToken,
		Expires:  userData.AccessToken.ExpiresIn,
		HTTPOnly: true,
	}

	c.Cookie(&cookieAccessToken)

	// set refresh token cookie
	cookieRefreshToken := fiber.Cookie{
		Name:     "refresh-token",
		Value:    userData.AccessToken.RefreshToken,
		Expires:  userData.AccessToken.ExpiresIn,
		HTTPOnly: true,
	}

	c.Cookie(&cookieRefreshToken)

	// return success response
	return c.Status(fiber.StatusOK).JSON(response.ApiResponse{
		Status:  response.SuccessStatus,
		Message: "Token refreshed successfully",
		Data:    userData.User,
	})

}
