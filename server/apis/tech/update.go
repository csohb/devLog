package tech

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UpdateTechRequest struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Percentage int    `json:"percentage"`
}

type ServiceUpdateTech struct {
	*context.Context
	req UpdateTechRequest
}

func (app *ServiceUpdateTech) Service() *api_context.CommonResponse {

	tid, err := strconv.Atoi(app.req.ID)
	if err != nil {
		app.Log.WithError(err).Error("not valid tech id")
		return api_context.FailureJSON(http.StatusBadRequest, "유효한 tech id가 아닙니다.")
	}
	tb := database.TBTech{
		Model: gorm.Model{
			ID: uint(tid),
		},
		Name:       app.req.Name,
		Percentage: app.req.Percentage,
	}

	if err = tb.Update(app.DB); err != nil {
		app.Log.WithError(err).Error("cannot update tech info")
		return api_context.FailureJSON(http.StatusInternalServerError, "데이터 업데이트 실패.")
	}

	return api_context.SuccessJSON(nil)
}

func (app *ServiceUpdateTech) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceUpdateTech) ApiName() interface{} {
	return "tech-update"
}

func (app ServiceUpdateTech) IsRequiredAuth() bool {
	return true
}

func ProcessUpdateTech(c echo.Context) error {
	return api_context.Process(&ServiceUpdateTech{
		Context: c.(*context.Context),
	})
}
