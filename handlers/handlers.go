package handlers

import "github.com/gofiber/fiber/v2"

func Home(ctx *fiber.Ctx) error {
	return ctx.SendString("product feedback")
}