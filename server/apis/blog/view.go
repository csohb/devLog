package blog

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UpdateViewRequest struct {
	BlogID string `json:"id"`
	Count  int    `json:"count"`
}

type UpdateViewResponse struct {
	View int `json:"view"`
}

type ServiceUpdateView struct {
	*context.Context
	req UpdateViewRequest
}

func (app *ServiceUpdateView) Service() *api_context.CommonResponse {
	tb := database.TBBlog{}
	if err := tb.UpdateView(app.DB, app.req.BlogID, app.req.Count); err != nil {
		app.Log.Errorf("update view err : %+v", app.req.BlogID)
		return api_context.FailureJSON(http.StatusInternalServerError, "update view err")
	}
	if err := tb.Get(app.DB, app.req.BlogID); err != nil {
		app.Log.Errorf("get blog data err : %+v", app.req.BlogID)
		return api_context.FailureJSON(http.StatusInternalServerError, "get heart err")
	}

	resp := UpdateViewResponse{View: tb.View}
	return api_context.SuccessJSON(&resp)
}

func (app *ServiceUpdateView) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceUpdateView) ApiName() interface{} {
	return "blog-heart"
}

func (app ServiceUpdateView) IsRequiredAuth() bool {
	return false
}

func ProcessUpdateView(c echo.Context) error {
	return api_context.Process(&ServiceUpdateView{
		Context: c.(*context.Context),
	})
}
