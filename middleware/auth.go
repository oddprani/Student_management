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
	user := &database.User{}
	session := &database.Session{}

	if err := database.DBInstance.Where("token = ?", token).First(session).Error; err != nil {
		return c.Next()
	}

	if err := database.DBInstance.Where("id = ?", session.UserID).First(user).Error; err != nil {
		return c.Next()
	}

	c.Locals("user", user)

	return c.Next()
}
