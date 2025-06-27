package router

import (
	"app/internal/controllers"
	"app/internal/services"

	"github.com/gofiber/fiber/v2"
)

func HealthCheckRoutes(app fiber.Router, h services.HealthCheckService) {
	healthCheckController := controllers.NewHealthCheckController(h)

	healthCheck := app.Group("/health-check")
	healthCheck.Get("/", healthCheckController.Check)
}
