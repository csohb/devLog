package mainpage

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Profile struct {
	UserID   string `json:"user_id"`
	Image    string `json:"image"`
	NickName string `json:"nick_name"`
	Dev      string `json:"dev"`
	Intro    string `json:"intro"`
}

type GetProfileListResponse struct {
	Profiles []Profile `json:"profiles"`
}

type ServiceProfileList struct {
	*context.Context
}

func (app *ServiceProfileList) Service() *api_context.CommonResponse {
	tb := database.TBUser{}
	list, err := tb.GetIntroduceForMainPage(app.DB)
	if err != nil {
		app.Log.Errorf("get introduce for mainpage err")
		return api_context.FailureJSON(http.StatusInternalServerError, "introduce 정보 불러오기 실패")
	}

	resp := GetProfileListResponse{}
	resp.Profiles = make([]Profile, len(list))
	for i, v := range list {
		var dev string
		switch v.Introduce.Developer {
		case "F":
			dev = "Frontend Developer"
		case "B":
			dev = "Backend Developer"
		}
		resp.Profiles[i] = Profile{
			UserID:   v.ID,
			Image:    v.Introduce.ProfileUrl,
			NickName: v.NickName,
			Dev:      dev,
			Intro:    v.Introduce.Intro,
		}
	}

	return api_context.SuccessJSON(&resp)
}

func (app *ServiceProfileList) GetRequestData() interface{} {
	return nil
}

func (app ServiceProfileList) ApiName() interface{} {
	return "get-profile-main"
}

func (app ServiceProfileList) IsRequiredAuth() bool {
	return false
}

func ProcessProfileList(c echo.Context) error {
	return api_context.Process(&ServiceProfileList{Context: c.(*context.Context)})
}
