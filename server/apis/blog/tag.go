package blog

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"fmt"
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

type TagBlog struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Date        string   `json:"date"`
	Writer      string   `json:"writer"`
	View        int      `json:"view"`
	Heart       int      `json:"heart"`
	Tags        []string `json:"tags"`
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

	list, total, err := tb.SearchTags(app.DB, app.req.Tag, app.req.Page-1, app.req.Count)
	if err != nil {
		logrus.WithError(err).Error("search with tags err : ", err)
		return api_context.FailureJSON(http.StatusInternalServerError, "find with tags err")
	}

	fmt.Println("list : ", list)
	resp.List = make([]Blog, len(list))
	for i, v := range list {
		tags := make([]string, len(v.Tags))
		for j, l := range v.Tags {
			tags[j] = l.Tag
		}
		resp.List[i] = Blog{
			ID:          strconv.Itoa(int(v.ID)),
			Title:       v.Title,
			Description: v.Description,
			Date:        v.CreatedAt.Format("2006-01-02"),
			Writer:      v.Writer,
			View:        v.View,
			Heart:       v.Heart,
			Tags:        tags,
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

func ProcessSearchWithTag(c echo.Context) error {
	return api_context.Process(&ServiceSearchWithTags{
		Context: c.(*context.Context),
	})
}
