package AdsControllers

import (
	AdsServices "backend-app/api/advertisements/services"

	"github.com/gofiber/fiber/v2"
)

func GetAdsList(c *fiber.Ctx) error {
	result, err := AdsServices.GetAdsList()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success",
		"result": result,
	})
}
