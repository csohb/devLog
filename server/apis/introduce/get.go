package introduce

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

type GetIntroduceRequest struct {
	ID string `param:"id"`
}

type Profi struct {
	Name      string `json:"name"`
	NickName  string `json:"nick_name"`
	Developer string `json:"developer"`
	Img       string `json:"img"`
	Addr      string `json:"addr"`
	Email     string `json:"email"`
	Birthday  string `json:"birthday"`
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

type Skill struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Percentage int    `json:"percentage"`
}

type Keyword struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetIntroduceResponse struct {
	Profile  Profi     `json:"profile"`
	Careers  []Career  `json:"careers"`
	Project  []Project `json:"project"`
	Keywords []Keyword `json:"keywords"`
	Skills   []Skill   `json:"skills"`
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

	fmt.Println("name :", strings.ReplaceAll(ret.Name, " ", ""))

	profi := Profi{
		Name:      strings.ReplaceAll(ret.Name, " ", ""),
		NickName:  ret.Introduce.UserID,
		Developer: devType,
		Img:       ret.Introduce.ProfileUrl,
		Addr:      ret.Addr,
		Email:     ret.Email,
		Birthday:  ret.Birthday.Format("2006-01-02"),
	}
	careerList := make([]Career, len(ret.Career))
	for i, v := range ret.Career {
		careerList[i] = Career{
			Company:   v.CompanyName,
			StartDate: v.StartDate.Format("2006-01-02"),
			EndDate:   v.EndDate.Format("2006-01-02"),
			JobTitle:  v.JobTitle,
			JobDetail: v.JobDetail,
		}
	}

	projectList := make([]Project, len(ret.Project))
	for i, v := range ret.Project {
		projectList[i] = Project{
			ID:          string(v.ID),
			Name:        v.Name,
			IsPersonal:  v.IsPersonal,
			StartDate:   v.StartDate.Format("2006-01-02"),
			EndDate:     v.EndDate.Format("2006-01-02"),
			Description: v.Description,
		}
	}

	var skillsNum int

	keywordList := make([]Keyword, len(ret.Tech))
	for i, v := range ret.Tech {
		if v.Percentage != 0 {
			skillsNum++
		}
		keywordList[i] = Keyword{
			ID:   strconv.Itoa(int(v.ID)),
			Name: v.Name,
		}
	}

	skillsList := make([]Skill, skillsNum)
	for i, v := range ret.Tech {
		if v.Percentage != 0 {
			skillsList[i] = Skill{
				ID:         strconv.Itoa(int(v.ID)),
				Name:       v.Name,
				Percentage: v.Percentage,
			}
		}
	}

	resp.Profile = profi
	resp.Careers = careerList
	resp.Project = projectList
	resp.Keywords = keywordList
	resp.Skills = skillsList

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
