package router

import (
	"app/internal/handler"
	s "app/internal/services"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router, a s.AuthService, u s.UserService,
	t s.TokenService) {
	authHandler := handler.NewAuthHandler(a, u, t)
	authPath := app.Group("/auth")

	authPath.Post("/login", authHandler.Login)
	authPath.Post("/register", authHandler.Register)
	authPath.Post("/refresh-token", authHandler.RefreshTokens)
	// v1 := app.Group("/v1")
	// v1.Post("/login", loginHandler)
	// v1.Post("/register", registerHandler)
	// v1.Post("/refresh-token", refreshTokenHandler)
	// v1.Post("/logout", logoutHandler)
}
