package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

// TracingMiddleware logs request duration and details
func TracingMiddleware(logger zerolog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			duration := time.Since(start)
			status := c.Response().Status
			method := c.Request().Method
			path := c.Request().URL.Path
			requestID := c.Get("request_id")

			logger.Info().
				Str("method", method).
				Str("path", path).
				Int("status", status).
				Dur("duration", duration).
				Interface("request_id", requestID).
				Msg("request completed")

			return err
		}
	}
}
