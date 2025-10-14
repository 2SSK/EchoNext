package middleware

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// StaticAppMiddleware configures static file serving for the Next.js app
func StaticAppMiddleware() echo.MiddlewareFunc {
	return middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "app/out",
		Index:  "index.html",
		Browse: false,
		HTML5:  true,
		Skipper: func(c echo.Context) bool {
			// Skip static middleware for API routes
			return strings.HasPrefix(c.Request().URL.Path, "/api")
		},
	})
}

// StaticAppFallbackHandler provides fallback for SPA routes
func StaticAppFallbackHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL.Path

		// Skip API routes
		if strings.HasPrefix(path, "/api") {
			return echo.NewHTTPError(http.StatusNotFound, "API endpoint not found")
		}

		// Check if it's a static asset request
		if strings.Contains(path, ".") {
			ext := filepath.Ext(path)
			if ext == ".js" || ext == ".css" || ext == ".ico" || ext == ".png" || ext == ".jpg" || ext == ".svg" {
				return echo.NewHTTPError(http.StatusNotFound, "Static file not found")
			}
		}

		// Serve index.html for all other routes (SPA routing)
		return c.File("app/out/index.html")
	}
}
