package story

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type GetStoryRequest struct {
	ID string `param:"id"`
}

type GetStoryResponse struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	Writer      string `json:"writer"`
	View        int    `json:"view"`
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

	resp := GetStoryResponse{
		ID:          strconv.Itoa(int(tb.ID)),
		CreatedAt:   tb.CreatedAt.Format("2006-01-02"),
		Title:       tb.Title,
		Content:     tb.Content,
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
