package login

import (
	"devLog/common/api_context"
	"devLog/server/apis/auth"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RequestLogin struct {
	UserID   string `json:"id"`
	Password string `json:"password"`
}

type ServiceLogin struct {
	*context.Context
	req RequestLogin
}

func (app *ServiceLogin) Service() *api_context.CommonResponse {

	// id find

	// pwd check

	// auth info create

	_, err := auth.CreateSession(app.Context, app.req.UserID)
	if err != nil {
		app.Log.WithError(err).Error("session create failed.")
		return api_context.FailureJSON(http.StatusInternalServerError, "세션 생성 실패")
	}

	return api_context.SuccessJSON("login success")
}

func (app *ServiceLogin) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceLogin) ApiName() interface{} {
	return "devLog-login"
}

func (app ServiceLogin) IsRequiredAuth() bool {
	return false
}

func ProcessLogin(c echo.Context) error {
	return api_context.Process(&ServiceLogin{
		Context: c.(*context.Context),
	})
}
