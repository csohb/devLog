package blog

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type SaveBlogRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	Images      []string `json:"images"`
	Writer      string   `json:"writer"`
	Tags        []string `json:"tags"`
}

type ServiceSaveBlog struct {
	*context.Context
	req SaveBlogRequest
}

func (app *ServiceSaveBlog) Service() *api_context.CommonResponse {
	bytes, err := json.Marshal(app.req.Images)
	if err != nil {
		app.Log.WithError(err).Error("json marshal error")
		return api_context.FailureJSON(http.StatusBadRequest, "cannot marshal images url")
	}

	userId := app.AuthInfo.GetUserID(app.Context)
	fmt.Println("userId : ", userId)

	tb := database.TBBlog{
		Model:       gorm.Model{CreatedAt: time.Now()},
		Title:       app.req.Title,
		Content:     app.req.Content,
		Image:       string(bytes),
		Description: app.req.Description,
		Writer:      app.AuthInfo.GetUserID(app.Context),
		Tags:        make([]database.TBTag, len(app.req.Tags)),
	}
	for i, v := range app.req.Tags {
		tb.Tags[i] = database.TBTag{
			Tag: v,
		}
	}

	err = app.DB.Create(&tb).Error
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
