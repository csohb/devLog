package blog

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

type GetBlogRequest struct {
	ID string `param:"id"`
}

type GetBlogResponse struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	Images      []string `json:"images"`
	Date        string   `json:"date"`
	Writer      string   `json:"writer"`
	View        int      `json:"view"`
	Heart       int      `json:"heart"`
	Tags        []string `json:"tags"`
}

type ServiceGetBlog struct {
	*context.Context
	req GetBlogRequest
}

func (app *ServiceGetBlog) Service() *api_context.CommonResponse {

	userID := app.AuthInfo.UserId
	logrus.Info("userID : ", userID)

	tb := database.TBBlog{}
	if err := tb.Get(app.DB, app.req.ID); err != nil {
		app.Log.Errorf("blog get err - blogID : %s", app.req.ID)
		return api_context.FailureJSON(http.StatusInternalServerError, "블로그 상세 정보 불러오기 실패")
	}

	images := make([]string, 10)

	if err := json.Unmarshal([]byte(tb.Image), &images); err != nil {
		app.Log.WithError(err).Error("cannot unmarshal images")
		return api_context.FailureJSON(http.StatusInternalServerError, "get db data err")
	}

	resp := GetBlogResponse{
		ID:          strconv.Itoa(int(tb.ID)),
		Title:       tb.Title,
		Description: tb.Description,
		Content:     tb.Content,
		Images:      images,
		Writer:      tb.Writer,
		Date:        tb.UpdatedAt.Format("2006-01-02"),
		View:        tb.View,
		Heart:       tb.Heart,
		Tags:        make([]string, len(tb.Tags)),
	}

	for i, v := range tb.Tags {
		resp.Tags[i] = v.Tag
	}

	return api_context.SuccessJSON(&resp)

}

func (app *ServiceGetBlog) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceGetBlog) ApiName() interface{} {
	return "get-blog"
}

func (app ServiceGetBlog) IsRequiredAuth() bool {
	return false
}

func ProcessGetBlog(c echo.Context) error {
	return api_context.Process(&ServiceGetBlog{Context: c.(*context.Context)})
}
