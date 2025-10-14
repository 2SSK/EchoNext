package sqlerr

import (
	"github.com/2SSK/EchoNext/internal/errs"
	"github.com/labstack/echo/v4"
)

// HandleSQLError handles SQL errors in Echo context
func HandleSQLError(c echo.Context, err error) error {
	apiErr := MapSQLError(err)
	if apiErr != nil {
		return c.JSON(apiErr.Status, apiErr)
	}
	return c.JSON(500, errs.InternalServer("Unknown error"))
}
