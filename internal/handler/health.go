package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/2SSK/EchoNext/internal/database"
	"github.com/labstack/echo/v4"
)

// HealthCheck handles GET /health
func HealthCheck(db *database.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check database connectivity
		ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
		defer cancel()

		if err := db.Pool.Ping(ctx); err != nil {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{
				"status":   "unhealthy",
				"database": "down",
				"error":    err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"status":    "healthy",
			"database":  "up",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	}
}
