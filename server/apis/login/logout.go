package login

import (
	"devLog/common/api_context"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ServiceLogout struct {
	*context.Context
}

func (app *ServiceLogout) Service() *api_context.CommonResponse {
	if err := app.AuthInfo.ExpireCookie(app.Context); err != nil {
		logrus.WithError(err).Error("Expire Cookie")
		return api_context.FailureJSON(http.StatusBadRequest, "logout error")
	}
	return nil
}

func (app *ServiceLogout) GetRequestData() interface{} {
	return nil
}

func (app ServiceLogout) ApiName() interface{} {
	return "logout"
}

func (app ServiceLogout) IsRequiredAuth() bool {
	return false
}

func ProcessLogout(c echo.Context) error {
	return api_context.Process(&ServiceLogout{c.(*context.Context)})
}
