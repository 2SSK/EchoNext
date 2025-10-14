package router

import (
	"github.com/2SSK/EchoNext/internal/database"
	mw "github.com/2SSK/EchoNext/internal/middleware"
	"github.com/2SSK/EchoNext/internal/routes"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, db *database.Database) {
	// Setup API routes first
	api := e.Group("/api")
	v1 := api.Group("/v1")
	routes.SetupRoutes(v1, db)

	// Configure static file serving for the Next.js app
	e.Use(mw.StaticAppMiddleware())

	// Fallback handler for SPA routes
	e.GET("/*", mw.StaticAppFallbackHandler())
}
