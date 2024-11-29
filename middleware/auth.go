package middleware

import (
	"studentmanager/database"

	"github.com/gofiber/fiber/v2"
)

func AuthenticationMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		return c.Next()
	}

	// Check session in the database and retrieve the user
	var user database.User

	if err := database.DBInstance.Where("token = ?", token).Preload("Sessions").First(&user).Error; err != nil {
		return c.Next()
	}

	c.Locals("user", user)

	return c.Next()
}
