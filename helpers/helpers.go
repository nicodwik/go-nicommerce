package helpers

import (
	"net/http"

	"github.com/labstack/echo"
)

var c echo.Context

func ServerErrorResponse(message string) error {

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"message": message,
	})
}
