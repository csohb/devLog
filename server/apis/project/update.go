package project

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"devLog/server/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UpdateProjectRequest struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	IsPersonal  bool     `json:"is_personal"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	Description string   `json:"description"`
	Stack       []string `json:"stack"`
}

type ServiceUpdateProject struct {
	*context.Context
	req UpdateProjectRequest
}

func (app *ServiceUpdateProject) Service() *api_context.CommonResponse {
	startDate, err := utils.StringToTime(app.req.StartDate)
	if err != nil {
		return api_context.FailureJSON(http.StatusBadRequest, "not valid startDate")
	}
	endDate, err := utils.StringToTime(app.req.EndDate)
	if err != nil {
		return api_context.FailureJSON(http.StatusBadRequest, "not valid endDate")
	}

	pid, err := strconv.Atoi(app.req.ID)
	if err != nil {
		logrus.WithError(err).Error("strconv failed.")
		return api_context.FailureJSON(http.StatusBadRequest, "not valid project ID")
	}

	stack := make([]database.TBStack, len(app.req.Stack))
	for i, v := range app.req.Stack {
		stack[i] = database.TBStack{
			Name: v,
		}
	}

	tb := database.TBProject{
		Model: gorm.Model{
			ID: uint(pid),
		},
		Name:        app.req.Name,
		IsPersonal:  app.req.IsPersonal,
		StartDate:   startDate,
		EndDate:     endDate,
		Description: app.req.Description,
	}

	err = tb.Update(app.DB, stack)
	if err != nil {
		logrus.WithError(err).Error("update project err")
		return api_context.InternalServerErrorJSON("project db update err")
	}

	return api_context.SuccessJSON(nil)
}

func (app *ServiceUpdateProject) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceUpdateProject) ApiName() interface{} {
	return "project-update"
}

func (app ServiceUpdateProject) IsRequiredAuth() bool {
	return true
}

func ProcessUpdateProject(c echo.Context) error {
	return api_context.Process(&ServiceUpdateProject{
		Context: c.(*context.Context),
	})
}
