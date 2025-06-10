package main

import (
	"github.com/gofiber/fiber/v2"
	"product_service/Routes"
)

func main() {
	app := fiber.New()
	Routes.Setup(app)
	app.Listen(":8001")
}
