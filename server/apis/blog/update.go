package blog

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type UpdateBlogRequest struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Images      []string `json:"images"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type ServiceUpdateBlog struct {
	*context.Context
	req UpdateBlogRequest
}

func (app *ServiceUpdateBlog) Service() *api_context.CommonResponse {
	bytes, err := json.Marshal(app.req.Images)
	if err != nil {
		app.Log.WithError(err).Error("json marshal error")
		return api_context.FailureJSON(http.StatusBadRequest, "cannot marshal images url")
	}

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
			ID:        uint(id),
			UpdatedAt: time.Now(),
		},
		Title:       app.req.Title,
		Description: app.req.Description,
		Content:     app.req.Content,
		Image:       string(bytes),
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
