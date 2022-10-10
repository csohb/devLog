package story

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UpdateStoryRequest struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	Content     string   `json:"content"`
	Images      []string `json:"images"`
	Description string   `json:"description"`
}

type ServiceUpdateStory struct {
	*context.Context
	req UpdateStoryRequest
}

func (app *ServiceUpdateStory) Service() *api_context.CommonResponse {

	sid, err := strconv.Atoi(app.req.ID)
	if err != nil {
		logrus.WithError(err).Error("sid convert err")
		return api_context.FailureJSON(http.StatusBadRequest, "not valid sid")
	}

	bytes, err := json.Marshal(app.req.Images)
	if err != nil {
		app.Log.WithError(err).Error("json marshal error")
		return api_context.FailureJSON(http.StatusBadRequest, "cannot marshal images url")
	}

	tb := database.TBStory{
		Model: gorm.Model{
			ID: uint(sid),
		},
		Title:       app.req.Title,
		Type:        app.req.Type,
		Content:     app.req.Content,
		Image:       string(bytes),
		Description: app.req.Description,
	}

	if err = tb.Update(app.DB); err != nil {
		logrus.WithError(err).Error("update story err")
		return api_context.FailureJSON(http.StatusInternalServerError, "update story err")
	}

	return api_context.SuccessJSON(nil)
}

func (app *ServiceUpdateStory) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceUpdateStory) ApiName() interface{} {
	return "story-update"
}

func (app ServiceUpdateStory) IsRequiredAuth() bool {
	return true
}

func ProcessUpdateStory(c echo.Context) error {
	return api_context.Process(&ServiceUpdateStory{
		Context: c.(*context.Context),
	})
}
