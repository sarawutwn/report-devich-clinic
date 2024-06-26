package AdvertisementRouter

import (
	AdsControllers "backend-app/api/advertisements/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	app := router.Group("promotion")
	app.Get("/get-all", AdsControllers.GetAdsList)
}
