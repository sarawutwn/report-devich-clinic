package CustomerRouter

import (
	CustomerControllers "backend-app/api/customer/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	app := router.Group("customer")
	app.Post("/create", CustomerControllers.CustomerNew)
	app.Patch("/update/status-reading/:customer_id", CustomerControllers.UpdateStatusReading)
	app.Put("/update/status-is-active", CustomerControllers.UpdateStatusIsActive)
	app.Get("/get-all", CustomerControllers.GetCustomer)
}
