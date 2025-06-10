package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"project_backend/Database"
	"project_backend/Models"
	"time"
)

var jwtSecret = []byte("your_secret_key")

func Register(c *fiber.Ctx) error {
	var input Models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	input.Password = string(hash)

	if err := Database.GormDB.Create(&input).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to register"})
	}

	return c.JSON(fiber.Map{"message": "Registered successfully"})
}

func Login(c *fiber.Ctx) error {
	var input Models.User
	var user Models.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := Database.GormDB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Wrong password"})
	}

	// JWT creation
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, _ := token.SignedString(jwtSecret)

	return c.JSON(fiber.Map{"token": tokenString})
}

func ValidateToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
	}

	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return c.JSON(fiber.Map{"user_id": claims["user_id"]})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token", "details": err.Error()})
	}
}

func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Successfully logged out"})
}
