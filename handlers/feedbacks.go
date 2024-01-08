package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/DelaRich/product-feedback-go/database"
	"github.com/DelaRich/product-feedback-go/helpers"
	"github.com/DelaRich/product-feedback-go/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	// Set status, created and updated time
	newFeedback.Status = "planned"
	newFeedback.CreatedAt = time.Now()
	newFeedback.UpdatedAt = time.Now()

	collection := database.GetCollection("feedbacks")

	// Check for existing feedback
	existingFeedback := models.Feedback{}
	ress := collection.FindOne(context.Background(), bson.M{"title": newFeedback.Title}).Decode(&existingFeedback)
	if ress == nil {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "Feedback with the same title already exists",
			"success": false,
		})

	}

	_, err := collection.InsertOne(context.Background(), newFeedback)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create feedback",
			"success": false,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Feedback added successfully",
		"success": true,
	})
}


func EditFeedback(ctx *fiber.Ctx) error {
	feedback := new(models.Feedback)


	if err := ctx.BodyParser(feedback); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// Validate titlte field
	if !helpers.IsValidInput(feedback.Title) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid title",
			"success": false,
		})
	}

	// Check for empty feedback title
	if feedback.Title == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Feedback title cannot be empty",
			"success": false,
		})
	}


	collection := database.GetCollection("feedbacks")
	filter := bson.M{"_id": feedback.Id}
	
	update := bson.M{"$set": bson.M{"title": feedback.Title, "category": feedback.Category, "details": feedback.Details, "status": feedback.Status, "updatedAt": time.Now()}}
	fmt.Println(filter)
res, err := collection.UpdateOne(context.Background(), filter, update)

if err != nil {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "Error updating feedback",
		"success": false,
	})
}

if res.ModifiedCount == 0 {
    return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
        "message": "Feedback not found",
        "success": false,
    })
}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Feedback updated successfully",
		"success": true,
	})
}

func GetFeedbacks(ctx *fiber.Ctx) error {
	collection := database.GetCollection("feedbacks")
	options := options.Find().SetSort(bson.D{{Key: "updatedAt", Value: -1}})
	cursor, err := collection.Find(context.Background(), bson.M{}, options)
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