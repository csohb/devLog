package api_context

import "github.com/labstack/echo/v4"

type AuthInterface interface {
	ErrorHandler() error
	Check() error
	Create(c echo.Context, data interface{}) error
}
