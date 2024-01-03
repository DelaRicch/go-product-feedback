package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2")

func main() {
	app := fiber.New()

	app.Use(cors.New())

	// setupRoutes
	setupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}