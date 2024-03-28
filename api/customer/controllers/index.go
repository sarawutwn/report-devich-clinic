package CustomerControllers

import (
	CustomerSchema "backend-app/api/customer/schema"
	CustomerServices "backend-app/api/customer/services"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

func CustomerNew(c *fiber.Ctx) error {
	body := new(CustomerSchema.CreateNewCustomer)
	c.BodyParser(&body)
	err := Validator.Struct(body)
	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field()+" is "+err.Tag())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "validate error",
			"message": errors,
		})
	}
	err = CustomerServices.CreateCustomer(body.CustomerName, body.AdsID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":  "success",
		"message": "Add new customer success.",
	})
}

func UpdateStatusReading(c *fiber.Ctx) error {
	result, err := CustomerServices.UpdateStatusReading(c.Params("customer_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":  "success",
		"result":  result,
		"message": "Update customer status successfully.",
	})
}

func UpdateStatusIsActive(c *fiber.Ctx) error {
	body := new(CustomerSchema.UpdateStatusIsActive)
	c.BodyParser(&body)
	err := Validator.Struct(body)
	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field()+" is "+err.Tag())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "validate error",
			"message": errors,
		})
	}
	result, err := CustomerServices.UpdateIsActive(body.CustomerID, body.IsActive)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":  "success",
		"result":  result,
		"message": "Update customer status successfully.",
	})
}

func GetCustomer(c *fiber.Ctx) error {
	result, err := CustomerServices.GetCustomer()
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
