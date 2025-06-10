package Routes

import (
	"github.com/gofiber/fiber/v2"
	Controllers "project_backend/controllers"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Post("/register", Controllers.Register)
	auth.Post("/login", Controllers.Login)
	auth.Get("/validate", Controllers.ValidateToken)
	auth.Post("/logout", Controllers.Logout)
}
