package main

import (
	"github.com/gofiber/fiber/v2"
	"project_backend/Database"
)

func main() {

	Database.InitDB()
	defer Database.DB.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Fiber + PostgreSQL + Viper works!")
	})

	app.Listen(":3000")
}
