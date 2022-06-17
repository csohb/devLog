package introduce

import (
	"devLog/common/api_context"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
)

type GetIntroduceRequest struct {
	UserID string `query:"id"`
}

type Career struct {
	CompanyName string `json:"company_name"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	JobTitle    string `json:"job_title"`
	JobDetail   string `json:"job_detail"`
}

type Project struct {
	Name        string  `json:"name"`
	IsPersonal  bool    `json:"is_personal"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	Description string  `json:"description"`
	Stack       []Stack `json:"stack"`
}

type Stack struct {
	Name string `json:"name"`
}

type GetIntroduceResponse struct {
	UserID     string    `json:"user_id"`
	Developer  string    `json:"developer"`
	Intro      string    `json:"intro"`
	ProfileUrl string    `json:"profile_url"`
	Career     []Career  `json:"career"`
	Projects   []Project `json:"projects"`
}

type ServiceGetIntroduce struct {
	*context.Context
	req GetIntroduceRequest
}

func (app *ServiceGetIntroduce) Service() *api_context.CommonResponse {
	//TODO implement me
	panic("implement me")
}

func (app *ServiceGetIntroduce) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceGetIntroduce) ApiName() interface{} {
	return "get-introduce"
}

func (app ServiceGetIntroduce) IsRequiredAuth() bool {
	return false
}

func ProcessGetIntroduce(c echo.Context) error {
	return api_context.Process(&ServiceGetIntroduce{Context: c.(*context.Context)})
}
