package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

// GlobalMiddlewares returns a list of all global middlewares
func GlobalMiddlewares(logger zerolog.Logger) []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		RequestIDMiddleware(),
		RateLimitMiddleware(1.0, 5, func(c echo.Context) bool {
			return !strings.HasPrefix(c.Request().URL.Path, "/api")
		}), // Skip rate limit for non-API paths
		TracingMiddleware(logger),
	}
}
