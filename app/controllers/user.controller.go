package controllers

import (
	"inventory/app/dtos"
	"inventory/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetUser(ctx *fiber.Ctx) error {
	product, err := services.GetUser(ctx.Params("id"))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(product)
}

func GetUsers(ctx *fiber.Ctx) error {
	products, err := services.GetUsers()
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(products)
}

func ChangePassword(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(string)

	var changePasswordDto dtos.ChangePasswordDto
	if err := ctx.BodyParser(&changePasswordDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := services.ChangePassword(userID, &changePasswordDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Password changed successfully"})
}
