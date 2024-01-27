package user

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/officemaid/app-api/app/data/response"
	userService "github.com/officemaid/app-api/app/services/user"
)

type GetUserRequest struct {
	Id string `json:"id"`
}

func GetUser(c *fiber.Ctx) error {

	log.Println("Get User")
	log.Println(c.Params("id"))

	user, err := userService.GetUser(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.ApiResponse{
			Status:  response.ErrorStatus,
			Message: "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.ApiResponse{
		Status:  response.SuccessStatus,
		Message: "",
		Data:    user,
	})
}

func GetUserList() string {
	return "Hello World"
}
