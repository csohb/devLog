package login

import (
	"devLog/common/api_context"
	"devLog/server/apis/auth"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ServiceLogout struct {
	*context.Context
}

func (app *ServiceLogout) Service() *api_context.CommonResponse {
	/*if err := app.AuthInfo.ExpireCookie(app.Context); err != nil {
		logrus.WithError(err).Error("Expire Cookie")
		return api_context.FailureJSON(http.StatusBadRequest, "logout error")
	}*/
	if err := auth.ExpireSession(app.Context); err != nil {
		return api_context.FailureJSON(http.StatusInternalServerError, "logout error")
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
