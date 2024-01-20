package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/officemaid/app-api/routes"
)

func main() {
	app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// 	AllowOrigins:     "*",
	// }))

	routes.Setup(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}

	// app.Get("/", func(c *fiber.Ctx) error {

	// 	password := "123"
	// 	dbHash := "$2a$12$lZRJf/dTrnSxJ.itxDaprePEm8a3TVwoQh0T3px.7uoIyadoAO0EG"

	// 	match := CheckPasswordHash(password, dbHash)

	// 	fmt.Println("Password:", password)
	// 	fmt.Println("Hash:    ", dbHash)
	// 	fmt.Println("Match:   ", match)

	// 	return c.SendString("Hello, World!2asasa22344")

	// })

	app.Listen(":3000")
}

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }
