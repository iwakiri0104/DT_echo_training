package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// index
func Index(c echo.Context) error {

	return c.Render(http.StatusOK, "index", nil)
}
