package main

import (
	GoCache "backend-app/cache/go-cache"

	"backend-app/database"
	"backend-app/router"

	Config "backend-app/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	GoCache.InitCache()
	app := fiber.New(fiber.Config{
		// Prefork: true,
	})
	// c := cron.New()
	// c.AddFunc("@every 110s", func() {
	// 	fmt.Println(time.Now())
	// })
	// c.Start()
	app.Use(cors.New(Config.CorsConfigDefault))
	router.SetupRoutes(app)
	app.Listen(":10000")

}
