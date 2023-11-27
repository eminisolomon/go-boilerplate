package controllers

import (
	"inventory/app/dtos"
	"inventory/app/models"
	"inventory/app/services"

	"github.com/gofiber/fiber/v2"
)

func Signup(ctx *fiber.Ctx) error {
	var newUser *models.User
	if err := ctx.BodyParser(&newUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := services.Signup(newUser)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}

func Signin(ctx *fiber.Ctx) error {
	var loginDto *dtos.LoginDto
	if err := ctx.BodyParser(&loginDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, accessToken, refreshToken, err := services.Signin(loginDto)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	response := fiber.Map{
		"user":         user,
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}

	return ctx.JSON(response)
}
