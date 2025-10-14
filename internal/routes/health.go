package routes

import (
	"github.com/2SSK/EchoNext/internal/database"
	"github.com/2SSK/EchoNext/internal/handler"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Group, db *database.Database) {
	e.GET("/health", handler.HealthCheck(db))
}
