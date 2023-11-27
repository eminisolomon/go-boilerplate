package routes

import (
	"inventory/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupStaffRoutes(api fiber.Router) {
	apiProduct := api.Group("/staff")
	apiProduct.Get("", controllers.GetStaffs)
	apiProduct.Post("", controllers.AddStaff)
	apiProduct.Post("/:staff_id", controllers.UpdateStaff)
	apiProduct.Get("/:staff_id", controllers.GetStaff)
	apiProduct.Delete("/:staff_id", controllers.DeleteStaff)
}
