package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func signIn(c *fiber.Ctx) error {
	fmt.Println("Hello World!")

	return c.JSON(fiber.Map{"token": "123"})
}

func signOut(c *fiber.Ctx) error {

}

func refreshToken(c *fiber.Ctx) error {

}

func me(c *fiber.Ctx) error {

}
