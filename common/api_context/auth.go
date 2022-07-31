package api_context

import "github.com/labstack/echo/v4"

type AuthInterface interface {
	ErrorHandler() error
	ParseAuthorization(c echo.Context) *CommonResponse
	Check() error
	Create(c echo.Context, data interface{}) error
}
