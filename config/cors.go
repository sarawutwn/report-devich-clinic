package config

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var CorsConfigDefault = cors.Config{
	Next: nil,
	// AllowOrigins: "*",
	AllowOrigins:     "http://localhost:5173, http://127.0.0.1:5173, https://react-devich-clinic.onrender.com, https://second-react-devich-clinic.onrender.com",
	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	AllowHeaders:     "",
	AllowCredentials: false,
	ExposeHeaders:    "",
	MaxAge:           30000,
}
