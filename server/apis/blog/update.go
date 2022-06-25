package blog

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UpdateBlogRequest struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type ServiceUpdateBlog struct {
	*context.Context
	req UpdateBlogRequest
}

func (app *ServiceUpdateBlog) Service() *api_context.CommonResponse {
	id, err := strconv.Atoi(app.req.ID)
	if err != nil {
		app.Log.Error("convert blog id to int err : ", app.req.ID)
		return api_context.FailureJSON(http.StatusBadRequest, "wrong blog id")
	}

	tags := make([]database.TBTag, len(app.req.Tags))
	for i, v := range app.req.Tags {
		tags[i] = database.TBTag{Tag: v}
	}

	tb := database.TBBlog{
		Model: gorm.Model{
			ID: uint(id),
		},
		Title:   app.req.Title,
		Content: app.req.Content,
	}

	if err = tb.Update(app.DB, tags); err != nil {
		app.Log.Error("update blog err : ", tb)
		return api_context.FailureJSON(http.StatusInternalServerError, "db update err")
	}

	return api_context.SuccessJSON(nil)
}

func (app *ServiceUpdateBlog) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceUpdateBlog) ApiName() interface{} {
	return "blog-update"
}

func (app ServiceUpdateBlog) IsRequiredAuth() bool {
	return false
}

func ProcessUpdateBlog(c echo.Context) error {
	return api_context.Process(&ServiceUpdateBlog{
		Context: c.(*context.Context),
	})
}
