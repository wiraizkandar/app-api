package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/officemaid/app-api/app/handlers/auth"
	"github.com/officemaid/app-api/app/handlers/user"
	"github.com/officemaid/app-api/app/middleware"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	// Version 1
	v1 := api.Group("/v1")

	// Auth Routes start
	authGroup := v1.Group("/auth")
	authGroup.Post("/authenticate", auth.Authenticate)
	authGroup.Post("/refresh", auth.RefreshToken)
	// Auth Routes End

	// User Routes start
	userGroup := v1.Group("/user", middleware.AuthMiddleware)
	userGroup.Get("/:id", user.GetUser)
	// User Routes end

	fmt.Println("Route Loaded")
}
