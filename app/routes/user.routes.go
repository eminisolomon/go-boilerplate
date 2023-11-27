package routes

import (
	"inventory/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(api fiber.Router) {
	apiProduct := api.Group("/users")
	apiProduct.Get("", controllers.GetUsers)
	apiProduct.Get("/change-password", controllers.ChangePassword)
	apiProduct.Get("/:id", controllers.GetUser)
}
