package blog

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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

	tb := database.TBBlog{}
	resp := SearchWithTagsResponse{}

	list, total, err := tb.SearchTags(app.DB, app.req.Tag, app.req.Page, app.req.Count)
	if err != nil {
		logrus.WithError(err).Error("search with tags err : ", err)
		return api_context.FailureJSON(http.StatusInternalServerError, "find with tags err")
	}

	resp.List = make([]Blog, len(list))
	for i, v := range list {
		tags := make([]string, len(v.Tags))
		for j, l := range v.Tags {
			tags[j] = l.Tag
		}
		resp.List[i] = Blog{
			ID:      strconv.Itoa(int(v.ID)),
			Title:   v.Title,
			Content: v.Content,
			Writer:  v.Writer,
			View:    v.View,
			Heart:   v.Heart,
			Tags:    tags,
		}
	}

	resp.Total = int(total)

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