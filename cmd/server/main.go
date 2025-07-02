package main

import (
	"app/internal/config"
	m "app/internal/middleware"
	"app/internal/router"
	"app/internal/utils"
	"app/pkg/database/ent"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	db "app/pkg/database"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := fiber.New(config.FiberConfig())
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(m.RecoverConfig())
	// Initialize database client
	client := db.DBConnect()
	defer func(Ent *ent.Client) {
		err := Ent.Close()
		if err != nil {

		}
	}(client.Ent)
	defer func(client *db.DBClient) {
		if err := client.Ent.Close(); err != nil {
			utils.Log.Errorf("Error closing Ent connection: %v", err)
		}
		if err := client.Redis.Close(); err != nil {
			utils.Log.Errorf("Error closing Redis connection: %v", err)
		}
	}(client)
	// Middleware
	router.SetupRoutes(app, client, ctx)
	app.Use(utils.NotFoundHandler)
	// Routes
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	// Start server
	if config.IsProd {
		serverErrors := make(chan error, 1)
		go startServer(app, fmt.Sprintf("%s:%d", config.AppHost, config.AppPort), serverErrors)
		handleGracefulShutdown(ctx, app, serverErrors) // Graceful shutdown, not recommended for Development
	} else {
		utils.Log.Info("Starting server in development mode...")
		if err := app.Listen(fmt.Sprintf("%s:%d", config.AppHost, config.AppPort)); err != nil {
			panic(err)
		}
	}

}

func startServer(app *fiber.App, address string, errs chan<- error) {
	if err := app.Listen(address); err != nil {
		errs <- fmt.Errorf("error starting server: %w", err)
	}
}
func handleGracefulShutdown(ctx context.Context, app *fiber.App, serverErrors <-chan error) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		utils.Log.Fatalf("Server error: %v", err)
	case <-quit:
		utils.Log.Info("Shutting down server...")
		if err := app.Shutdown(); err != nil {
			utils.Log.Fatalf("Error during server shutdown: %v", err)
		}
	case <-ctx.Done():
		utils.Log.Info("Server exiting due to context cancellation")
	}

	utils.Log.Info("Server exited")
}
