package main

import (
	"log"

	"github.com/DelaRich/product-feedback-go/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	  // Establish database connection
	  if err := database.ConnectDb(); err != nil {
        log.Fatal(err)
    }
    defer database.DisconnectDb()

	app := fiber.New()
	app.Use(cors.New())

	// setupRoutes
	setupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}