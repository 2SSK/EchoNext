package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// RequestIDMiddleware adds a unique request ID to each request
func RequestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := uuid.New().String()
			c.Set("request_id", requestID)
			c.Response().Header().Set("X-Request-ID", requestID)
			return next(c)
		}
	}
}
