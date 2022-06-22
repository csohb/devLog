package blog

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UpdateHeartRequest struct {
	BlogID string `json:"id"`
	Count  int    `json:"count"`
}

type UpdateHeartResponse struct {
	Heart int `json:"heart"`
}

type ServiceUpdateHeart struct {
	*context.Context
	req UpdateHeartRequest
}

func (app *ServiceUpdateHeart) Service() *api_context.CommonResponse {
	tb := database.TBBlog{}
	if err := tb.UpdateHeart(app.DB, app.req.BlogID, app.req.Count); err != nil {
		app.Log.Errorf("update heart err : %+v", app.req.BlogID)
		return api_context.FailureJSON(http.StatusInternalServerError, "update heart err")
	}
	if err := tb.Get(app.DB, app.req.BlogID); err != nil {
		app.Log.Errorf("get blog data err : %+v", app.req.BlogID)
		return api_context.FailureJSON(http.StatusInternalServerError, "get heart err")
	}

	resp := UpdateHeartResponse{Heart: tb.Heart}
	return api_context.SuccessJSON(&resp)
}

func (app *ServiceUpdateHeart) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceUpdateHeart) ApiName() interface{} {
	return "blog-heart"
}

func (app ServiceUpdateHeart) IsRequiredAuth() bool {
	return false
}

func ProcessUpdateHeart(c echo.Context) error {
	return api_context.Process(&ServiceUpdateHeart{
		Context: c.(*context.Context),
	})
}
