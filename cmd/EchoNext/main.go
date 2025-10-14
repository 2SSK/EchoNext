package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/2SSK/EchoNext/internal/config"
	"github.com/2SSK/EchoNext/internal/database"
	"github.com/2SSK/EchoNext/internal/logger"
	"github.com/2SSK/EchoNext/internal/router"
	"github.com/2SSK/EchoNext/internal/server"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Get environment
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// Initialize logger
	appLogger := logger.NewLogger(env)

	// Load database configuration
	dsn, err := config.LoadDatabaseURL()
	if err != nil {
		appLogger.Fatal().Err(err).Msg("Failed to load database configuration")
	}

	// Initialize database connection
	db, err := database.New(dsn, &appLogger)
	if err != nil {
		appLogger.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer db.Close()

	// Load server configuration
	serverCfg, err := config.LoadServerConfig()
	if err != nil {
		appLogger.Fatal().Err(err).Msg("Failed to load server configuration")
	}

	// Initialize Echo instance
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Add middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	// Setup routes
	router.SetupRoutes(e, db)

	// Create HTTP server
	srv := server.New(e, serverCfg)

	// Start server in a goroutine
	go func() {
		appLogger.Info().Str("port", serverCfg.Port).Msg("Starting EchoNext server")
		if err := srv.ListenAndServe(); err != nil && err.Error() != "http: Server closed" {
			appLogger.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	appLogger.Info().Msg("Shutting down server...")

	// Give outstanding requests 30 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		appLogger.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	appLogger.Info().Msg("Server exited")
}
