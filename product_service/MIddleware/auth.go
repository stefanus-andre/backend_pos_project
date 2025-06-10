package MIddleware

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Tidak ada token"})
	}

	req, err := http.NewRequest("GET", "http://localhost:8000/validate", nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": err.Error()})
	}
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Error": "Unauthorized"})
	}
	return c.Next()
}
