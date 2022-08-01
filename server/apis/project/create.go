package project

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"devLog/server/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Project struct {
	Name        string   `json:"name"`
	IsPersonal  bool     `json:"is_personal"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	Description string   `json:"description"`
	Stack       []string `json:"stack"`
}

type CreateProjectRequest struct {
	ID       string    `json:"id"`
	Projects []Project `json:"projects"`
}

type ServiceCreateProject struct {
	*context.Context
	req CreateProjectRequest
}

func (app *ServiceCreateProject) Service() *api_context.CommonResponse {
	for _, v := range app.req.Projects {
		startTime, err := utils.StringToTime(v.StartDate)
		if err != nil {
			return api_context.FailureJSON(http.StatusBadRequest, "not valid startDate")
		}
		endTime, err := utils.StringToTime(v.EndDate)
		if err != nil {
			return api_context.FailureJSON(http.StatusBadRequest, "not valid endDate")
		}

		stack := make([]database.TBStack, len(v.Stack))
		for j, k := range v.Stack {
			stack[j] = database.TBStack{
				Name: k,
			}
		}
		tb := database.TBProject{
			UserID:      app.req.ID,
			Name:        v.Name,
			IsPersonal:  v.IsPersonal,
			StartDate:   startTime,
			EndDate:     endTime,
			Description: v.Description,
			Stack:       stack,
		}

		fmt.Println("tbProject : ", tb)

		if err = tb.Save(app.DB); err != nil {
			return api_context.FailureJSON(http.StatusInternalServerError, "db insert error")
		}
	}

	return api_context.SuccessJSON(nil)
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
