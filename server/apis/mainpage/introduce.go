package mainpage

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type IntroduceRequest struct {
	ID string `param:"id"`
	//ID string `json:"id"`
}

type Profi struct {
	NickName string `json:"nick_name"`
}

type Career struct {
	Company   string `json:"company"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	JobTitle  string `json:"job_title"`
	JobDetail string `json:"job_detail"`
}

type IntroduceResponse struct {
	Profile Profi    `json:"profile"`
	Careers []Career `json:"careers"`
}

type ServiceIntroduce struct {
	*context.Context
	req IntroduceRequest
}

func (app *ServiceIntroduce) Service() *api_context.CommonResponse {
	resp := IntroduceResponse{}
	fmt.Printf("user id : %s", app.req.ID)
	tb := database.TBUser{}
	err := tb.GetCareer(app.DB, app.req.ID)
	if err != nil {
		logrus.WithError(err).Error("get profi err")
		return api_context.FailureJSON(http.StatusInternalServerError, "프로필 조회에 실패하였습니다.")
	}

	profi := Profi{
		NickName: tb.NickName,
	}
	careerList := make([]Career, len(tb.Career))
	for i, v := range tb.Career {
		careerList[i] = Career{
			Company:   v.CompanyName,
			StartDate: v.StartDate.Format("2006-01-02"),
			EndDate:   v.EndDate.Format("2006-01-02"),
			JobTitle:  v.JobTitle,
			JobDetail: v.JobDetail,
		}
	}

	resp.Profile = profi
	resp.Careers = careerList

	return api_context.SuccessJSON(&resp)
}

func (app *ServiceIntroduce) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceIntroduce) ApiName() interface{} {
	return "introduce"
}

func (app ServiceIntroduce) IsRequiredAuth() bool {
	return false
}

func ProcessIntroduce(c echo.Context) error {
	return api_context.Process(&ServiceIntroduce{Context: c.(*context.Context)})
}
