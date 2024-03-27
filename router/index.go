package router

import (
	ReportRouter "backend-app/api/report"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api", logger.New())
	ReportRouter.SetupRoutes(api)
}
