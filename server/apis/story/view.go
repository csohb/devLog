package story

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ViewStoryRequest struct {
	ID string `param:"id"`
}

type ServiceViewStory struct {
	*context.Context
	req ViewStoryRequest
}

func (app *ServiceViewStory) Service() *api_context.CommonResponse {
	tb := database.TBStory{}
	if err := tb.UpdateView(app.DB, app.req.ID); err != nil {
		logrus.WithError(err).Error("update view err")
		return api_context.FailureJSON(http.StatusInternalServerError, "update view err")
	}

	return api_context.SuccessJSON(nil)
}

func (app *ServiceViewStory) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceViewStory) ApiName() interface{} {
	return "story-view"
}

func (app ServiceViewStory) IsRequiredAuth() bool {
	return false
}

func ProcessViewStory(c echo.Context) error {
	return api_context.Process(&ServiceViewStory{
		Context: c.(*context.Context),
	})
}
