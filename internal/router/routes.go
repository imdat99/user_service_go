package router

import (
	"app/internal/services"
	"app/pkg/database/ent"
	"context"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *ent.Client, ctx context.Context) {
	v1 := app.Group("/v1")
	healthCheckService := services.NewHealthCheckService(db, &ctx)
	HealthCheckRoutes(v1, healthCheckService)
}
