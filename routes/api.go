package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/officemaid/app-api/app/handlers/auth"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	// Version 1

	v1 := api.Group("/v1")

	// Auth Routes start
	authGroup := v1.Group("/auth")

	authGroup.Post("/sign-in", auth.SignIn)
	// Auth Routes End

	fmt.Println("Route Loaded2")
}
