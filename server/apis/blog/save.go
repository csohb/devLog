package blog

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SaveBlogRequest struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Writer  string   `json:"writer"`
	Tags    []string `json:"tags"`
}

type ServiceSaveBlog struct {
	*context.Context
	req SaveBlogRequest
}

func (app *ServiceSaveBlog) Service() *api_context.CommonResponse {
	tb := database.TBBlog{
		Title:   app.req.Title,
		Content: app.req.Content,
		Writer:  app.req.Writer,
		Tags:    make([]database.TBTag, len(app.req.Tags)),
	}
	for i, v := range app.req.Tags {
		tb.Tags[i] = database.TBTag{
			Tag: v,
		}
	}

	err := app.DB.Create(&tb).Error
	if err != nil {
		app.Log.Errorf("Create blog err : %+v", err)
		return api_context.FailureJSON(http.StatusInternalServerError, "db create err")
	}

	return api_context.SuccessJSON(nil)
}

func (app *ServiceSaveBlog) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceSaveBlog) ApiName() interface{} {
	return "save-blog"
}

func (app ServiceSaveBlog) IsRequiredAuth() bool {
	return true
}

func ProcessSaveBlog(c echo.Context) error {
	return api_context.Process(&ServiceSaveBlog{Context: c.(*context.Context)})
}
