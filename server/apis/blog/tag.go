package blog

import (
	"devLog/common/api_context"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
)

type SearchWithTagsRequest struct {
	Tag   string `query:"tag"`
	Page  int    `query:"page"`
	Count int    `query:"count"`
}

type SearchWithTagsResponse struct {
	List  []Blog `json:"list"`
	Total int    `json:"total"`
}

type ServiceSearchWithTags struct {
	*context.Context
	req SearchWithTagsRequest
}

func (app *ServiceSearchWithTags) Service() *api_context.CommonResponse {

	resp := SearchWithTagsResponse{}

	return api_context.SuccessJSON(&resp)
}

func (app *ServiceSearchWithTags) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceSearchWithTags) ApiName() interface{} {
	return "blog-search-tags"
}

func (app ServiceSearchWithTags) IsRequiredAuth() bool {
	return false
}

func (ServiceSearchWithTags) Process(c echo.Context) error {
	return api_context.Process(&ServiceSearchWithTags{
		Context: c.(*context.Context),
	})
}