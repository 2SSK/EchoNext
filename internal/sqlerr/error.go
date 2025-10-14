package sqlerr

import (
	"errors"
	"strings"

	"github.com/2SSK/EchoNext/internal/errs"
	"github.com/jackc/pgx/v5/pgconn"
)

// MapSQLError maps SQL errors to API errors
func MapSQLError(err error) *errs.APIError {
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505": // unique_violation
			return errs.Conflict("Resource already exists")
		case "23503": // foreign_key_violation
			return errs.BadRequest("Invalid reference")
		case "23502": // not_null_violation
			return errs.BadRequest("Required field is missing")
		case "23514": // check_violation
			return errs.BadRequest("Data validation failed")
		case "42P01": // undefined_table
			return errs.InternalServer("Database schema issue")
		case "08003": // connection_does_not_exist
			return errs.InternalServer("Database connection lost")
		case "08006": // connection_failure
			return errs.InternalServer("Database connection failed")
		default:
			return errs.InternalServer("Database error: " + pgErr.Message)
		}
	}

	// Check for common error messages
	errMsg := strings.ToLower(err.Error())
	if strings.Contains(errMsg, "no rows") {
		return errs.NotFound("Resource not found")
	}
	if strings.Contains(errMsg, "connection refused") {
		return errs.InternalServer("Database unavailable")
	}
	if strings.Contains(errMsg, "timeout") {
		return errs.InternalServer("Database timeout")
	}

	// For other errors, return generic internal error
	return errs.InternalServer("Database operation failed")
}
