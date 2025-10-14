package errs

import "net/http"

// Common error codes
const (
	ErrBadRequest     = "BAD_REQUEST"
	ErrUnauthorized   = "UNAUTHORIZED"
	ErrForbidden      = "FORBIDDEN"
	ErrNotFound       = "NOT_FOUND"
	ErrConflict       = "CONFLICT"
	ErrInternalServer = "INTERNAL_SERVER_ERROR"
)

// Helper functions for common errors
func BadRequest(message string) *APIError {
	return NewAPIError(http.StatusBadRequest, ErrBadRequest, message)
}

func Unauthorized(message string) *APIError {
	return NewAPIError(http.StatusUnauthorized, ErrUnauthorized, message)
}

func Forbidden(message string) *APIError {
	return NewAPIError(http.StatusForbidden, ErrForbidden, message)
}

func NotFound(message string) *APIError {
	return NewAPIError(http.StatusNotFound, ErrNotFound, message)
}

func Conflict(message string) *APIError {
	return NewAPIError(http.StatusConflict, ErrConflict, message)
}

func InternalServer(message string) *APIError {
	return NewAPIError(http.StatusInternalServerError, ErrInternalServer, message)
}
