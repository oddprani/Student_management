package controllers

import (
	"studentmanager/config"
	"studentmanager/database"
	"studentmanager/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	user := &database.User{}

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorResponse{
			Message:   "Invalid request body",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	if user.Username == "" || user.Password == "" || user.Type == "" {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorResponse{
			Message:   "Invalid request body",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&config.ErrorResponse{
			Message:   "Could not hash password",
			ErrorCode: fiber.StatusInternalServerError,
		})
	}

	user.Password = string(hashedPassword)

	if err := database.DBInstance.Create(user).Error; err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(&config.ErrorResponse{
			Message:   "Could not create user",
			ErrorCode: fiber.StatusNotAcceptable,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func Login(c *fiber.Ctx) error {

	user := &database.User{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&config.ErrorResponse{
			Message:   "Invalid request body",
			ErrorCode: fiber.StatusBadRequest,
		})
	}

	requestPassword := user.Password

	// Check if user exists
	if err := database.DBInstance.Where("username = ?", user.Username).First(user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&config.ErrorResponse{
			Message:   "Invalid credentials",
			ErrorCode: fiber.StatusUnauthorized,
		})
	}

	// Check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestPassword)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&config.ErrorResponse{
			Message:   "Invalid credentials",
			ErrorCode: fiber.StatusUnauthorized,
		})
	}

	// Create session
	session := &database.Session{
		Token:  utils.GenerateToken(),
		UserID: user.ID,
	}

	if err := database.DBInstance.Create(session).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&config.ErrorResponse{
			Message:   "Could not create session",
			ErrorCode: fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(session)
}

// This is a test route to check if the authentication middleware works
func AuthRoute(c *fiber.Ctx) error {
	user := c.Locals("user")

	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&config.ErrorResponse{
			Message:   "Unauthorized",
			ErrorCode: fiber.StatusUnauthorized,
		})
	}

	user = user.(*database.User)

	return c.Status(fiber.StatusOK).JSON(user)
}
