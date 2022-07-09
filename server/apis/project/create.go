package project

import (
	"devLog/common/api_context"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
)

type CreateProjectRequest struct {
}

type ServiceCreateProject struct {
	*context.Context
	req CreateProjectRequest
}

func (app *ServiceCreateProject) Service() *api_context.CommonResponse {
	//TODO implement me
	panic("implement me")
}

func (app *ServiceCreateProject) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceCreateProject) ApiName() interface{} {
	return "project-create"
}

func (app ServiceCreateProject) IsRequiredAuth() bool {
	return true
}

func ProcessCreateProject(c echo.Context) error {
	return api_context.Process(&ServiceCreateProject{
		Context: c.(*context.Context),
	})
}
