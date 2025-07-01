package router

import (
	"app/internal/handler"
	"app/internal/services"

	"github.com/gofiber/fiber/v2"
)

func HealthCheckRoutes(app fiber.Router, h services.HealthCheckService) {
	healthCheckController := handler.NewHealthCheckHandler(h)

	healthCheck := app.Group("/health-check")
	healthCheck.Get("/", healthCheckController.Check)
}
