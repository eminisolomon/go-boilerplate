package app

import (
	"fmt"

	cfg "inventory/app/core"
	db "inventory/app/database"
	"inventory/app/models"
	"inventory/app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run() {
	app := fiber.New()

	cfg.LoadConfig()
	config := cfg.GetConfig()

	db.ConnectPostgres()
	// db.PgDB.Migrator().DropTable(&user.User{})
	db.PgDB.AutoMigrate(&models.User{})
	db.ConnectMongo()
	db.ConnectRedis()

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	api := app.Group("/api")
	routes.SetupAuthRoutes(api)
	routes.SetupUserRoutes(api)
	routes.SetupStaffRoutes(api)

	port := fmt.Sprintf(":%s", config.Port)
	app.Listen(port)
}
