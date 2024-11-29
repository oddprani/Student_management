package router

import (
	"studentmanager/config"
	"studentmanager/controllers"

	"github.com/gofiber/fiber/v2"
)

func Initialise(router *fiber.App) {
	userRouter := router.Group("/user")
	userRouter.Get("/auth", controllers.AuthRoute)
	userRouter.Post("/register", controllers.Register)
	userRouter.Post("/login", controllers.Login)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(config.ErrorResponse{
			Message:   "Not found",
			ErrorCode: fiber.StatusNotFound,
		})
	})
}
