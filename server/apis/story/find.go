package story

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"fmt"
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
	Content     string `json:"content"`
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
		fmt.Println("v : ", v)
		resp.List[i] = Story{
			ID:          strconv.Itoa(int(v.ID)),
			CreatedAt:   v.CreatedAt.Format("2006-01-02"),
			Title:       v.Title,
			Content:     v.Content,
			Description: v.Description,
			Writer:      v.Writer,
			View:        v.View,
		}
	}

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
