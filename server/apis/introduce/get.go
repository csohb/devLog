package introduce

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type GetIntroduceRequest struct {
	ID string `param:"id"`
}

type Profi struct {
	NickName  string `json:"nick_name"`
	Developer string `json:"developer"`
	Img       string `json:"img"`
}

type Career struct {
	Company   string `json:"company"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	JobTitle  string `json:"job_title"`
	JobDetail string `json:"job_detail"`
}

type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	IsPersonal  bool   `json:"is_personal"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
}

type GetIntroduceResponse struct {
	Profile Profi     `json:"profile"`
	Careers []Career  `json:"careers"`
	Project []Project `json:"project"`
}

type ServiceIntroduce struct {
	*context.Context
	req GetIntroduceRequest
}

func (app *ServiceIntroduce) Service() *api_context.CommonResponse {
	resp := GetIntroduceResponse{}
	fmt.Printf("user id : %s", app.req.ID)
	tb := database.TBUser{}
	ret, err := tb.GetIntroduce(app.DB, app.req.ID)
	if err != nil {
		logrus.WithError(err).Error("get profi err")
		return api_context.FailureJSON(http.StatusInternalServerError, "프로필 조회에 실패하였습니다.")
	}

	var devType string
	if ret.Introduce.Developer == "F" {
		devType = "Frontend 개발자"
	} else if ret.Introduce.Developer == "B" {
		devType = "Backend 개발자"
	}
	profi := Profi{
		NickName:  ret.Introduce.UserID,
		Developer: devType,
		Img:       ret.Introduce.ProfileUrl,
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

	projectList := make([]Project, len(tb.Project))
	for i, v := range tb.Project {
		projectList[i] = Project{
			ID:          string(v.ID),
			Name:        v.Name,
			IsPersonal:  v.IsPersonal,
			StartDate:   v.StartDate.Format("2006-01-02"),
			EndDate:     v.EndDate.Format("2006-01-02"),
			Description: v.Description,
		}
	}

	resp.Profile = profi
	resp.Careers = careerList
	resp.Project = projectList

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

func ProcessGetIntroduce(c echo.Context) error {
	return api_context.Process(&ServiceIntroduce{Context: c.(*context.Context)})
}
