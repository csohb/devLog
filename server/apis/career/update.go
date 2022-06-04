package career

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type UpdateCareerRequest struct {
	ID        int    `json:"id"`
	Company   string `json:"company"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	JobTitle  string `json:"job_title"`
	JobDetail string `json:"job_detail"`
}

type ServiceUpdateCareer struct {
	*context.Context
	req UpdateCareerRequest
}

func (app ServiceUpdateCareer) setStringToTime(str string) (time.Time, error) {
	return time.Parse("2006-01-02", str)
}

func (app *ServiceUpdateCareer) Service() *api_context.CommonResponse {
	start, err := app.setStringToTime(app.req.StartDate)
	if err != nil {
		app.Log.Errorf("parse string to time failed. : %+v", app.req.StartDate)
	}
	end, err := app.setStringToTime(app.req.EndDate)
	if err != nil {
		app.Log.Errorf("parse string to time failed. : %+v", app.req.StartDate)
	}

	tb := database.TBCareer{
		Model:       gorm.Model{ID: uint(app.req.ID)},
		CompanyName: app.req.Company,
		StartDate:   start,
		EndDate:     end,
		JobTitle:    app.req.JobTitle,
		JobDetail:   app.req.JobDetail,
	}

	if err = tb.UpdateCareer(app.DB, tb); err != nil {
		app.Log.Errorf("update career failed : %+v", tb)
		return api_context.FailureJSON(http.StatusInternalServerError, "career 업데이트에 실패하였습니다.")
	}

	return api_context.SuccessJSON(nil)
}

func (app *ServiceUpdateCareer) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceUpdateCareer) ApiName() interface{} {
	return "update-career"
}

func (app ServiceUpdateCareer) IsRequiredAuth() bool {
	return true
}

func ProcessUpdateCareer(c echo.Context) error {
	return api_context.Process(&ServiceUpdateCareer{Context: c.(*context.Context)})
}
