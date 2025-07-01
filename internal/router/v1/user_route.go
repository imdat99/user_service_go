package router

import (
	"app/internal/handler"
	"app/internal/middleware"
	s "app/internal/services"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app fiber.Router, u s.UserService) {
	userHandler := handler.NewUserHandler(u)

	userPath := app.Group("/users")
	userPath.Use(middleware.Auth(u))

	userPath.Get("/me", userHandler.MyInfo)
}
