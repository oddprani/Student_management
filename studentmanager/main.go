package main

import (
	"log"
	"studentmanager/config"
	"studentmanager/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	if config.Debug {
		log.Printf("Running in debug mode on port %s", config.Port)
		log.Printf("Config: %+v", config)
	} else {
		log.Printf("Running in production mode on port %s", config.Port)
	}

	database.Initialize()

	app := fiber.New()

	app.Use(recover.New(recover.Config{
		EnableStackTrace: config.Debug,
	}))

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	app.Use(logger.New())
	app.Use(helmet.New())

	log.Fatal(app.Listen(":" + config.Port))
}
