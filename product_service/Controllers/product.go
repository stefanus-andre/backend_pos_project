package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"product_service/Database"
	"product_service/Models"
)

func GetProducts(c *fiber.Ctx) error {
	var products []Models.Product
	Database.GormDB.Find(&products)
	return c.Status(200).JSON(products)
}

func GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var product Models.Product
	if err := Database.GormDB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.Status(200).JSON(fiber.Map{"product": product})
}

func CreateProduct(c *fiber.Ctx) error {
	product := Models.Product{}
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse the JSON"})
	}
	Database.GormDB.Create(&product)
	return c.Status(200).JSON(fiber.Map{"product": product})
}
