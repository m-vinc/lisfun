package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrInternalServerError = echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	ErrNotFound            = echo.NewHTTPError(http.StatusNotFound, "not found")
)
