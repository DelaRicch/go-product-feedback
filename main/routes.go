package main

import (
	"github.com/DelaRich/product-feedback-go/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Get("/feedbacks", handlers.GetFeedbacks)
	app.Post("add-feedback", handlers.AddFeedback)
	app.Patch("/edit-feedback", handlers.EditFeedback)
	app.Delete("/delete-feedbacks", handlers.DeleteAllFeedbacks)
}