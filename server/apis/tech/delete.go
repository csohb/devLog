package tech

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DeleteTechRequest struct {
	ID string `param:"id"`
}

type ServiceDeleteTech struct {
	*context.Context
	req DeleteTechRequest
}

func (app *ServiceDeleteTech) Service() *api_context.CommonResponse {
	tb := database.TBTech{}
	if err := tb.Delete(app.DB, app.req.ID); err != nil {
		app.Log.WithError(err).Error("cannot delete tech info")
		return api_context.FailureJSON(http.StatusInternalServerError, "데이터 삭제에 실패하였습니다.")
	}

	return api_context.SuccessJSON(nil)
}

func (app *ServiceDeleteTech) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceDeleteTech) ApiName() interface{} {
	return "tech-delete"
}

func (app ServiceDeleteTech) IsRequiredAuth() bool {
	return true
}

func ProcessDeleteTech(c echo.Context) error {
	return api_context.Process(&ServiceDeleteTech{
		Context: c.(*context.Context),
	})
}
