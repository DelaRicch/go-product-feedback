package handlers

import (
	"context"
	"time"

	"github.com/DelaRich/product-feedback-go/database"
	"github.com/DelaRich/product-feedback-go/helpers"
	"github.com/DelaRich/product-feedback-go/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Home(ctx *fiber.Ctx) error {
	return ctx.SendString("product feedback")
}

func AddFeedback(ctx *fiber.Ctx) error {
	newFeedback := new(models.Feedback)

	if err := ctx.BodyParser(newFeedback); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// Validate titlte field
	if !helpers.IsValidInput(newFeedback.Title) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid title",
			"success": false,
		})
	}

	// Check for empty feedback title
	if newFeedback.Title == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Feedback title cannot be empty",
			"success": false,
		})
	}

	// Set created and updated time
	newFeedback.CreatedAt = time.Now()
	newFeedback.UpdatedAt = time.Now()

	collection := database.GetCollection("feedbacks")

	// Check for existing feedback
	existingFeedback := models.Feedback{}
	ress := collection.FindOne(context.Background(), bson.M{"title": newFeedback.Title}).Decode(&existingFeedback)
	if ress == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Feedback wit the same title already exists",
			"success": false,
		})

	}

	_, err := collection.InsertOne(context.Background(), newFeedback)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to insert user",
			"success": false,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Feedback added successfully",
		"success": true,
	})
}

func GetFeedbacks(ctx *fiber.Ctx) error {
	collection := database.GetCollection("feedbacks")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error fetching users",
			"success": false,
		})
	}

	defer cursor.Close(context.Background())

	feedbacks := []models.Feedback{}
	for cursor.Next(context.Background()) {
		var feedback models.Feedback
		if err := cursor.Decode(&feedback); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to decode feedback",
				"sucess":  false,
			})
		}
		feedbacks = append(feedbacks, feedback)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Successfully retrieved feedbacks",
		"success":    true,
		"feedbacks": feedbacks,
	})
}

func DeleteAllFeedbacks(ctx *fiber.Ctx) error {
	collection := database.GetCollection("feedbacks")
	_, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error deleting feedbacks",
			"success": false,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "All feedbacks deleted successfully",
		"success": true,
	})
}