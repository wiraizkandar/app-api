package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	authservice "github.com/officemaid/app-api/app/services/auth"
)

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignIn(c *fiber.Ctx) error {

	fmt.Println("Username")

	request := new(SignInRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	response := authservice.AuthenticateUser(request.Username, request.Password)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}