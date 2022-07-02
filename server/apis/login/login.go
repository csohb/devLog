package login

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/auth"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
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
	tb := database.TBUser{}
	if err := tb.CheckID(app.DB, app.req.UserID); err != nil {
		app.Log.WithError(err).Error("not exist id")
		return api_context.FailureJSON(http.StatusBadRequest, "존재하지 않는 아이디입니다.")
	}
	// pwd check
	logrus.Infof("req.ID : %+v", app.req.UserID)
	logrus.Infof("req.password : %+v", app.req.Password)

	if err := tb.CheckPwd(app.DB, app.req.UserID, app.req.Password); err != nil {
		app.Log.WithError(err).Error("not correct pwd")
		return api_context.FailureJSON(http.StatusBadRequest, "비밀번호가 일치하지 않습니다.")
	}
	// auth info create
	authInfo := context.AuthHandler{
		UserId:    app.req.UserID,
		LoginDate: time.Now(),
	}
	app.AuthInfo.Create(app.Context, authInfo)

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
