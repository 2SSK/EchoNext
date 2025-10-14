package middleware

import (
	"net"
	"sync"

	"github.com/2SSK/EchoNext/internal/errs"
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

// RateLimitMiddleware creates a rate limiting middleware
func RateLimitMiddleware(rps float64, burst int, skipper func(echo.Context) bool) echo.MiddlewareFunc {
	limiters := make(map[string]*rate.Limiter)
	mu := sync.Mutex{}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if skipper(c) {
				return next(c)
			}

			ip := getRealIP(c)

			mu.Lock()
			limiter, exists := limiters[ip]
			if !exists {
				limiter = rate.NewLimiter(rate.Limit(rps), burst)
				limiters[ip] = limiter
			}
			mu.Unlock()

			if !limiter.Allow() {
				return errs.NewAPIError(429, "RATE_LIMIT_EXCEEDED", "Too many requests", nil)
			}

			return next(c)
		}
	}
}

// getRealIP extracts the real IP address
func getRealIP(c echo.Context) string {
	// Check X-Forwarded-For header
	xff := c.Request().Header.Get("X-Forwarded-For")
	if xff != "" {
		// Take the first IP
		ip := net.ParseIP(xff)
		if ip != nil {
			return ip.String()
		}
	}

	// Check X-Real-IP header
	xri := c.Request().Header.Get("X-Real-IP")
	if xri != "" {
		ip := net.ParseIP(xri)
		if ip != nil {
			return ip.String()
		}
	}

	// Fallback to RemoteAddr
	ip, _, err := net.SplitHostPort(c.Request().RemoteAddr)
	if err != nil {
		return c.Request().RemoteAddr
	}
	return ip
}
