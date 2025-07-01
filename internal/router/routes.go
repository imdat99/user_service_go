package router

import (
	rv1 "app/internal/router/v1"
	"app/internal/services"
	"app/internal/validation"
	database "app/pkg/database"
	"context"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *database.DBClient, ctx context.Context) {
	validate := validation.Validator()
	api := app.Group("/api")
	v1 := api.Group("/v1")

	userService := services.NewUserService(db.Ent, validate)
	tokenService := services.NewTokenService(db, validate)
	authService := services.NewAuthService(db.Ent, validate, userService, tokenService)
	healthCheckService := services.NewHealthCheckService(db.Ent, &ctx)

	HealthCheckRoutes(api, healthCheckService)
	rv1.AuthRoutes(v1, authService, userService, tokenService)
	rv1.UserRoutes(v1, userService)
}
