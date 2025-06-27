package main

import (
	"app/internal/config"
	"app/internal/router"
	"app/internal/utils"
	"app/pkg/database/ent"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"entgo.io/ent/dialect"
	esql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewClient() (*ent.Client, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Tối ưu connection pool
	db.SetMaxOpenConns(25)                 // Max connections
	db.SetMaxIdleConns(10)                 // Max idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Connection lifetime
	db.SetConnMaxIdleTime(2 * time.Minute) // Idle timeout

	drv := esql.OpenDB(dialect.MySQL, db)
	return ent.NewClient(ent.Driver(drv)), err
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := fiber.New(config.FiberConfig())
	app.Use(cors.New())
	app.Use(logger.New())
	// Initialize database client
	client, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	// Middleware
	router.SetupRoutes(app, client, ctx)
	app.Use(utils.NotFoundHandler)
	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start server
	// if err := app.Listen(":3000"); err != nil {
	// 	panic(err)
	// }
	serverErrors := make(chan error, 1)
	go startServer(app, ":3000", serverErrors)
	handleGracefulShutdown(ctx, app, serverErrors) // Graceful shutdown, not recommended for Development
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
