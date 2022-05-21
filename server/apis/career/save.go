package career

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"devLog/server/apis/mainpage"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type SaveCareerRequest struct {
	ID     string            `json:"id"`
	Career []mainpage.Career `json:"career"`
}

type ServiceSaveCareer struct {
	*context.Context
	req SaveCareerRequest
}

func (app ServiceSaveCareer) setStringToTime(str string) (time.Time, error) {
	return time.Parse("2006-01-02", str)
}

func (app *ServiceSaveCareer) Service() *api_context.CommonResponse {

	careerList := make([]database.TBCareer, len(app.req.Career))
	for i, v := range app.req.Career {
		var startDate time.Time
		var endDate time.Time
		var err error
		if v.StartDate != "" {
			startDate, err = app.setStringToTime(v.StartDate)
			if err != nil {
				logrus.WithError(err).Error("parse time string")
				return api_context.FailureJSON(http.StatusInternalServerError, "time parse err")
			}
		}

		if v.EndDate != "" {
			endDate, err = app.setStringToTime(v.EndDate)
			if err != nil {
				logrus.WithError(err).Error("parse time string")
				return api_context.FailureJSON(http.StatusInternalServerError, "time parse err")
			}
		}
		careerList[i] = database.TBCareer{
			UserID:      app.req.ID,
			CompanyName: v.Company,
			StartDate:   startDate,
			EndDate:     endDate,
			JobTitle:    v.JobTitle,
			JobDetail:   v.JobDetail,
		}
	}

	if err := app.DB.Model(&database.TBCareer{}).Create(careerList).Error; err != nil {
		logrus.WithError(err).Error("Save Carrer err")
		return api_context.FailureJSON(http.StatusInternalServerError, "이력 저장에 실패 (db 오류)")
	}

	return api_context.SuccessJSON(nil)
}

func (app *ServiceSaveCareer) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceSaveCareer) ApiName() interface{} {
	return "save-career"
}

func (app ServiceSaveCareer) IsRequiredAuth() bool {
	return false
}

func (app ServiceSaveCareer) GetAuthHandler() api_context.AuthInterface {
	//TODO implement me
	panic("implement me")
}

func ProcessSaveCareer(c echo.Context) error {
	return api_context.Process(&ServiceSaveCareer{Context: c.(*context.Context)})
}
