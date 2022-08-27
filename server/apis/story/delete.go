package story

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type DeleteStoryRequest struct {
	ID string `param:"id"`
}

type ServiceDeleteStory struct {
	*context.Context
	req DeleteStoryRequest
}

func (app *ServiceDeleteStory) Service() *api_context.CommonResponse {
	tb := database.TBStory{}
	if err := tb.Delete(app.DB, app.req.ID); err != nil {
		logrus.WithError(err).Error("story delete err")
		return api_context.FailureJSON(http.StatusInternalServerError, "story delete db err")
	}
	return api_context.SuccessJSON(nil)
}

func (app *ServiceDeleteStory) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceDeleteStory) ApiName() interface{} {
	return "story-delete"
}

func (app ServiceDeleteStory) IsRequiredAuth() bool {
	return true
}

func ProcessDeleteStory(c echo.Context) error {
	return api_context.Process(&ServiceDeleteStory{
		Context: c.(*context.Context),
	})
}
