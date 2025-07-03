package router

import (
	// initialize the Swagger documentation
	_ "app/cmd/server/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerRoutes(v1 fiber.Router) {
	docs := v1.Group("/swagger")

	docs.Get("/*", swagger.HandlerDefault)
}
