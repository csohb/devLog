package introduce

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UpdateProfi struct {
	NickName string `json:"nick_name"`
	Img      string `json:"img"`
	Intro    string `json:"intro"`
	Email    string `json:"email"`
	Addr     string `json:"addr"`
}

type UpdateIntroduceRequest struct {
	UserID  string      `json:"user_id"`
	Profile UpdateProfi `json:"profile"`
}

type ServiceUpdateIntroduce struct {
	*context.Context
	req UpdateIntroduceRequest
}

func (app *ServiceUpdateIntroduce) Service() *api_context.CommonResponse {
	tb := database.TBUser{
		ID:    app.req.UserID,
		Email: app.req.Profile.Email,
		Addr:  app.req.Profile.Addr,
		Introduce: database.TBIntroduce{
			Intro:      app.req.Profile.Intro,
			ProfileUrl: app.req.Profile.Img,
		},
	}
	if err := tb.UpdateIntroduce(app.DB, app.req.UserID); err != nil {
		app.Log.Error("update introduce failed.")
		return api_context.FailureJSON(http.StatusInternalServerError, "update db err")
	}

	return api_context.SuccessJSON(nil)
}

func (app *ServiceUpdateIntroduce) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceUpdateIntroduce) ApiName() interface{} {
	return "introduce-update"
}

func (app ServiceUpdateIntroduce) IsRequiredAuth() bool {
	return true
}

func ProcessUpdateIntroduce(c echo.Context) error {
	return api_context.Process(&ServiceUpdateIntroduce{
		Context: c.(*context.Context),
	})
}
