package api_context

import "github.com/labstack/echo/v4"

type AuthInterface interface {
	ParseAuthorization(c echo.Context) *CommonResponse
	Check() error
	GetUserID(c echo.Context) string
	Create(c echo.Context, data interface{}) error
}
