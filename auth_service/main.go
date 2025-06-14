package main

import (
	"github.com/gofiber/fiber/v2"
	"project_backend/Database"
	"project_backend/Routes"
)

func main() {

	Database.InitDB()
	Database.Migrate()

	defer Database.DB.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Fiber + PostgreSQL + Viper works!")
	})

	Routes.AuthRoutes(app)

	app.Listen(":8000")
}
