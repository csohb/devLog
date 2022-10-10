package story

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type FindStoryRequest struct {
	Page  int `query:"page"`
	Count int `query:"count"`
}

type Story struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Writer      string `json:"writer"`
	View        int    `json:"view"`
}

type FindStoryResponse struct {
	List  []Story `json:"list"`
	Total int     `json:"total"`
}

type ServiceFindStory struct {
	*context.Context
	req FindStoryRequest
}

func (app *ServiceFindStory) Service() *api_context.CommonResponse {

	tb := database.TBStory{}
	ret, total, err := tb.Find(app.DB, app.req.Page-1, app.req.Count)
	if err != nil {
		logrus.WithError(err).Error("find story err")
		return api_context.FailureJSON(http.StatusInternalServerError, "find data err")
	}

	resp := FindStoryResponse{}
	resp.List = make([]Story, len(ret))
	for i, v := range ret {
		var work string
		if v.Type == "F" {
			work = "FrontEnd"
		} else if v.Type == "B" {
			work = "BackEnd"
		} else {
			work = ""
		}

		images := make([]string, 10)
		if len(v.Image) > 0 {
			if err := json.Unmarshal([]byte(v.Image), &images); err != nil {
				app.Log.WithError(err).Error("cannot unmarshal images : ", v.Image)
				return api_context.FailureJSON(http.StatusInternalServerError, "get db data err")
			}
		}

		resp.List[i] = Story{
			ID:          strconv.Itoa(int(v.ID)),
			CreatedAt:   v.CreatedAt.Format("2006-01-02"),
			Title:       v.Title,
			Type:        work,
			Content:     v.Content,
			Image:       images[0],
			Description: v.Description,
			Writer:      v.Writer,
			View:        v.View,
		}
	}

	//app.Log.Info("resp.List : ", resp.List)

	resp.Total = int(total)
	return api_context.SuccessJSON(&resp)
}

func (app *ServiceFindStory) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceFindStory) ApiName() interface{} {
	return "story-find"
}

func (app ServiceFindStory) IsRequiredAuth() bool {
	return false
}

func ProcessFindStory(c echo.Context) error {
	return api_context.Process(&ServiceFindStory{
		Context: c.(*context.Context),
	})
}
