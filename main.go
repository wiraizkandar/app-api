package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/officemaid/app-api/routes"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func main() {
	app := fiber.New()

	godotenv.Load()

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-store")
		return c.Next()
	})

	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// 	AllowOrigins:     "*",
	// }))

	// app.Get("/", func(c *fiber.Ctx) error {

	// 	// password := "123"
	// 	// dbHash := "$2a$12$lZRJf/dTrnSxJ.itxDaprePEm8a3TVwoQh0T3px.7uoIyadoAO0EG"

	// 	// match := CheckPasswordHash(password, dbHash)

	// 	// fmt.Println("Password:", password)
	// 	// fmt.Println("Hash:    ", dbHash)
	// 	// fmt.Println("Match:   ", match)

	// 	return c.SendString("Hello, World!2asasa22344")

	// })

	routes.Setup(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }
