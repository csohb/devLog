package introduce

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type GetIntroduceRequest struct {
	ID string `param:"id"`
}

type Profi struct {
	Name        string `json:"name"`
	NickName    string `json:"nick_name"`
	Developer   string `json:"developer"`
	Description string `json:"description"`
	Img         string `json:"img"`
	Addr        string `json:"addr"`
	Email       string `json:"email"`
	Birthday    string `json:"birthday"`
}

type CareerDB struct {
	ID        string    `json:"id"`
	Company   string    `json:"company"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	JobTitle  string    `json:"job_title"`
	JobDetail string    `json:"job_detail"`
}

type Career struct {
	ID        string `json:"id"`
	Company   string `json:"company"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	JobTitle  string `json:"job_title"`
	JobDetail string `json:"job_detail"`
}

type ProjectDB struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	IsPersonal  bool      `json:"is_personal"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
	Stack       []string  `json:"stack"`
}

type Project struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	IsPersonal  bool     `json:"is_personal"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	Description string   `json:"description"`
	Stack       []string `json:"stack"`
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

func sorttime(dat []ProjectDB) {
	for i := 0; i < len(dat); i++ {
		for j := i + 1; j < len(dat); j++ {
			if dat[i].StartDate.Unix() > dat[j].StartDate.Unix() {
				dat[i], dat[j] = dat[j], dat[i]
			}
		}
	}
}

func sortCareer(dat []CareerDB) {
	for i := 0; i < len(dat); i++ {
		for j := i + 1; j < len(dat); j++ {
			if dat[i].StartDate.Unix() > dat[j].StartDate.Unix() {
				dat[i], dat[j] = dat[j], dat[i]
			}
		}
	}
}

func (app *ServiceIntroduce) Service() *api_context.CommonResponse {
	resp := GetIntroduceResponse{}
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
		Name:        strings.ReplaceAll(ret.Name, " ", ""),
		NickName:    ret.Introduce.UserID,
		Developer:   devType,
		Img:         ret.Introduce.ProfileUrl,
		Description: ret.Introduce.Intro,
		Addr:        ret.Addr,
		Email:       ret.Email,
		Birthday:    ret.Birthday.Format("2006-01-02"),
	}
	careerList := make([]CareerDB, len(ret.Career))
	for i, v := range ret.Career {
		careerList[i] = CareerDB{
			ID:        strconv.Itoa(int(v.ID)),
			Company:   v.CompanyName,
			StartDate: v.StartDate,
			EndDate:   v.EndDate,
			JobTitle:  v.JobTitle,
			JobDetail: v.JobDetail,
		}
	}
	sortCareer(careerList)

	careerResp := make([]Career, len(careerList))
	for i, v := range careerList {
		careerResp[i] = Career{
			ID:        v.ID,
			Company:   v.Company,
			StartDate: v.StartDate.Format("2006-01-02"),
			EndDate:   v.EndDate.Format("2006-01-02"),
			JobTitle:  v.JobTitle,
			JobDetail: v.JobDetail,
		}
	}

	projectList := make([]ProjectDB, len(ret.Project))
	for i, v := range ret.Project {
		tb := database.TBProject{}
		tb.Get(app.DB, v.ID)
		stack := make([]string, len(tb.Stack))
		for j, k := range tb.Stack {
			stack[j] = k.Name
		}
		projectList[i] = ProjectDB{
			ID:          strconv.Itoa(int(v.ID)),
			Name:        v.Name,
			IsPersonal:  v.IsPersonal,
			StartDate:   v.StartDate,
			EndDate:     v.EndDate,
			Description: v.Description,
			Stack:       stack,
		}
	}

	sorttime(projectList)

	proResp := make([]Project, len(projectList))
	for i, v := range projectList {
		proResp[i] = Project{
			ID:          v.ID,
			Name:        v.Name,
			IsPersonal:  v.IsPersonal,
			StartDate:   v.StartDate.Format("2006-01-02"),
			EndDate:     v.EndDate.Format("2006-01-02"),
			Description: v.Description,
			Stack:       v.Stack,
		}
	}

	var skillsNum int

	app.Log.Info("tech", ret.Tech)

	keywordList := make([]Keyword, len(ret.Tech))
	for i, v := range ret.Tech {
		if v.Percentage > 0 {
			skillsNum++
		}
		keywordList[i] = Keyword{
			ID:   strconv.Itoa(int(v.ID)),
			Name: v.Name,
		}
	}

	app.Log.Info("keywords", keywordList)
	app.Log.Info("skillsNum:", skillsNum)

	skillsList := make([]Skill, skillsNum)
	var index int
	for _, v := range ret.Tech {
		if v.Percentage > 0 {
			index++
			skillsList[index-1] = Skill{
				ID:         strconv.Itoa(int(v.ID)),
				Name:       v.Name,
				Percentage: v.Percentage,
			}
		}
	}

	resp.Profile = profi
	resp.Careers = careerResp
	resp.Project = proResp
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
