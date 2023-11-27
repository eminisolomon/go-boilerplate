package controllers

import (
	"inventory/app/models"
	"inventory/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetStaffs(ctx *fiber.Ctx) error {
	staff, err := services.GetStaffs()
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(staff)
}

func GetStaff(ctx *fiber.Ctx) error {
	staff, err := services.GetStaff(ctx.Params("id"))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(staff)
}

func AddStaff(ctx *fiber.Ctx) error {
	var newStaff *models.User
	if err := ctx.BodyParser(&newStaff); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	staff, err := services.AddStaff(newStaff)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(staff)
}

func UpdateStaff(ctx *fiber.Ctx) error {
	var staff *models.User
	if err := ctx.BodyParser(&staff); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	staff, err := services.UpdateStaff(staff, ctx.Params("id"))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(staff)
}
func DeleteStaff(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	err := services.DeleteStaff(userId)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.SendStatus(204)
}
