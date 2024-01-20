package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")
	/*
	 * Version 1
	 */
	fmt.Println("Hello World!")
	v1 := api.Group("/v1")
	// v1 := api.Group("/v1", middlewares.RequestAuthentication)

	auth := v1.Group("/auth")
	auth.Post("/sign-in", handlers.signIn)
	// auth.Post("/sign-out", handlers.CreateUser)
	// auth.Post("/refresh-token", handlers.AuthenticateUser)
	// auth.Post("/me", handlers.AuthenticateUser)

	// cdn := app.Group("/cdn")
	// cdn.Use("/images", filesystem.New(
	// 	filesystem.Config{
	// 		Root: pkger.Dir("/public/images"),
	// 	},
	// ))
}
