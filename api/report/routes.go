package ReportRouter

import (
	ReportController "backend-app/api/report/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	app := router.Group("reports")
	app.Get("/dashboard", ReportController.ReportDashbord)
	app.Get("/report-date", ReportController.ReportDashbordOnDate)
	app.Get("/report-quota/:page_size/:page", ReportController.GetQuotaByStoreID)
	app.Get("/report-users/:page_size/:page", ReportController.GetUsersByStoreID)
	app.Get("/report-transection/:page_size/:page", ReportController.GetTransectionByStoreID)
	app.Get("/export-users", ReportController.ExportUsersByStoreID)
	app.Get("/export-quota", ReportController.ExportQuotaByStoreID)
	app.Get("/export-transection", ReportController.ExportTransectionByStoreID)

}
