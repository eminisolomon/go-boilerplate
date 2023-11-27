package routes

import (
	"inventory/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(api fiber.Router) {
	apiUser := api.Group("/auth")
	apiUser.Post("/register", controllers.Signup)
	apiUser.Post("/login", controllers.Signin)
}
