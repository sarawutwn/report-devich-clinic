package reportControllers

import (
	ReportServices "backend-app/api/report/services"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ReportDashbord(c *fiber.Ctx) error {
	result, err := ReportServices.ReportDashbord(c.Locals("store_id").(string))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "not found",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success",
		"result": result,
	})
}

func ReportDashbordOnDate(c *fiber.Ctx) error {
	date := strings.Split(c.Query("date"), " - ")
	// fmt.Println(date)
	result, err := ReportServices.ReportDashbordOnDate(c.Locals("store_id").(string), date[0], date[1])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "not found",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success",
		"result": result,
	})
}

func GetQuotaByStoreID(c *fiber.Ctx) error {
	date := strings.Split(c.Query("date"), " - ")
	pageSize, _ := strconv.Atoi(c.Params("page_size"))
	page, _ := strconv.Atoi(c.Params("page"))
	// fmt.Println(date)
	result, count, err := ReportServices.GetQuotaByStoreID(c.Locals("store_id").(string), date[0], date[1], pageSize, page)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "not found",
			"count":   0,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success",
		"count":  count,
		"result": result,
	})
}

func GetUsersByStoreID(c *fiber.Ctx) error {
	date := strings.Split(c.Query("date"), " - ")
	pageSize, _ := strconv.Atoi(c.Params("page_size"))
	page, _ := strconv.Atoi(c.Params("page"))
	// fmt.Println(date)
	result, count, err := ReportServices.GetUsersByStoreID(c.Locals("store_id").(string), date[0], date[1], pageSize, page)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "not found",
			"count":   0,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success",
		"count":  count,
		"result": result,
	})
}

func GetTransectionByStoreID(c *fiber.Ctx) error {
	date := strings.Split(c.Query("date"), " - ")
	pageSize, _ := strconv.Atoi(c.Params("page_size"))
	page, _ := strconv.Atoi(c.Params("page"))
	// fmt.Println(date)
	result, count, err := ReportServices.GetTransectionByStoreID(c.Locals("store_id").(string), date[0], date[1], pageSize, page)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "not found",
			"count":   0,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success",
		"count":  count,
		"result": result,
	})
}

func ExportUsersByStoreID(c *fiber.Ctx) error {
	date := strings.Split(c.Query("date"), " - ")
	// fmt.Println(date)
	result, err := ReportServices.ExportUsersByStoreID(c.Locals("store_id").(string), date[0], date[1])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "not found",
			"message": err.Error(),
		})
	}

	_, err = os.Stat(result)
	if os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).SendString("File not found")
	}

	c.Set("Content-Disposition", "attachment; filename=your_excel_file.xlsx")
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	return c.SendFile(result)
	// return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
	// 	"status": "success",
	// 	"result": result,
	// })
}

func ExportQuotaByStoreID(c *fiber.Ctx) error {
	date := strings.Split(c.Query("date"), " - ")
	// fmt.Println(date)
	result, err := ReportServices.ExportQuotaByStoreID(c.Locals("store_id").(string), date[0], date[1])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "not found",
			"message": err.Error(),
		})
	}

	_, err = os.Stat(result)
	if os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).SendString("File not found")
	}

	c.Set("Content-Disposition", "attachment; filename=your_excel_file.xlsx")
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	return c.SendFile(result)
	// return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
	// 	"status": "success",
	// 	"result": result,
	// })
}

func ExportTransectionByStoreID(c *fiber.Ctx) error {
	date := strings.Split(c.Query("date"), " - ")
	// fmt.Println(date)
	result, err := ReportServices.ExportTransectionByStoreID(c.Locals("store_id").(string), date[0], date[1])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "not found",
			"message": err.Error(),
		})
	}

	_, err = os.Stat(result)
	if os.IsNotExist(err) {
		return c.Status(fiber.StatusNotFound).SendString("File not found")
	}

	c.Set("Content-Disposition", "attachment; filename=your_excel_file.xlsx")
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	return c.SendFile(result)
	// return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
	// 	"status": "success",
	// 	"result": result,
	// })
}
