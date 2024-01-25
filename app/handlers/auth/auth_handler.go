package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/officemaid/app-api/app/data/response"
	authservice "github.com/officemaid/app-api/app/services/auth"
)

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/**
 * @api {post} /api/v1/auth/sign-in Sign In
 */
func Authenticate(c *fiber.Ctx) error {

	request := new(SignInRequest)

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

	// return success response
	return c.Status(fiber.StatusOK).JSON(response.ApiResponse{
		Status:  response.SuccessStatus,
		Message: "User authenticated successfully",
		Data:    userData,
	})

}
