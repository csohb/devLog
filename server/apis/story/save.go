package story

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SaveStoryRequest struct {
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	Content     string   `json:"content"`
	Images      []string `json:"images"`
	Description string   `json:"description"`
}

type ServiceSaveStory struct {
	*context.Context
	req SaveStoryRequest
}

func (app *ServiceSaveStory) Service() *api_context.CommonResponse {
	bytes, err := json.Marshal(app.req.Images)
	if err != nil {
		app.Log.WithError(err).Error("json marshal error")
		return api_context.FailureJSON(http.StatusBadRequest, "cannot marshal images url")
	}

	userId := app.AuthInfo.GetUserID(app.Context)
	fmt.Println("userId : ", userId)
	tb := database.TBStory{
		Title:       app.req.Title,
		Type:        app.req.Type,
		Content:     app.req.Content,
		Image:       string(bytes),
		Description: app.req.Description,
		Writer:      app.AuthInfo.GetUserID(app.Context),
	}

	if err := tb.Save(app.DB); err != nil {
		logrus.WithError(err).Error("save story err")
		return api_context.FailureJSON(http.StatusInternalServerError, "save story db err")
	}

	return api_context.SuccessJSON(nil)
}

func (app *ServiceSaveStory) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceSaveStory) ApiName() interface{} {
	return "story-save"
}

func (app ServiceSaveStory) IsRequiredAuth() bool {
	return true
}

func ProcessSaveStory(c echo.Context) error {
	return api_context.Process(&ServiceSaveStory{
		Context: c.(*context.Context),
	})
}
