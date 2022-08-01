package project

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type DeleteProjectRequest struct {
	ID string `param:"id"`
}

type ServiceDeleteProject struct {
	*context.Context
	req DeleteProjectRequest
}

func (app *ServiceDeleteProject) Service() *api_context.CommonResponse {
	tb := database.TBProject{}
	if err := tb.Delete(app.DB, app.req.ID); err != nil {
		logrus.WithError(err).Error("project delete failed.")
		return api_context.InternalServerErrorJSON("project delete failed.")
	}
	return api_context.SuccessJSON(nil)
}

func (app *ServiceDeleteProject) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceDeleteProject) ApiName() interface{} {
	return "project-delete"
}

func (app ServiceDeleteProject) IsRequiredAuth() bool {
	return true
}

func ProcessDeleteProject(c echo.Context) error {
	return api_context.Process(&ServiceDeleteProject{
		Context: c.(*context.Context),
	})
}
