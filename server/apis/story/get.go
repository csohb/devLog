package story

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type GetStoryRequest struct {
	ID string `param:"id"`
}

type GetStoryResponse struct {
	ID          string   `json:"id"`
	CreatedAt   string   `json:"created_at"`
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	Content     string   `json:"content"`
	Images      []string `json:"images"`
	Description string   `json:"description"`
	Writer      string   `json:"writer"`
	View        int      `json:"view"`
}

type ServiceGetStory struct {
	*context.Context
	req GetStoryRequest
}

func (app *ServiceGetStory) Service() *api_context.CommonResponse {
	tb := database.TBStory{}
	if err := tb.Get(app.DB, app.req.ID); err != nil {
		logrus.WithError(err).Error("get story err")
		return api_context.FailureJSON(http.StatusInternalServerError, "get db data err")
	}

	images := make([]string, 10)

	if err := json.Unmarshal([]byte(tb.Image), &images); err != nil {
		app.Log.WithError(err).Error("cannot unmarshal images")
		return api_context.FailureJSON(http.StatusInternalServerError, "get db data err")
	}

	resp := GetStoryResponse{
		ID:          strconv.Itoa(int(tb.ID)),
		CreatedAt:   tb.CreatedAt.Format("2006-01-02"),
		Title:       tb.Title,
		Type:        tb.Type,
		Content:     tb.Content,
		Images:      images,
		Description: tb.Description,
		Writer:      tb.Writer,
		View:        tb.View,
	}

	return api_context.SuccessJSON(&resp)

}

func (app *ServiceGetStory) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceGetStory) ApiName() interface{} {
	return "story-get"
}

func (app ServiceGetStory) IsRequiredAuth() bool {
	return false
}

func ProcessGetStory(c echo.Context) error {
	return api_context.Process(&ServiceGetStory{
		Context: c.(*context.Context),
	})
}
