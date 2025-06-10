package Routes

import (
	"github.com/gofiber/fiber/v2"
	"product_service/Controllers"
)

func Setup(app *fiber.App) {
	app.Get("/products", Controllers.GetProducts)
	app.Get("/products/:id", Controllers.GetProductByID)
	app.Post("/products", Controllers.CreateProduct)

}
