package blog

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type GetBlogRequest struct {
	ID string `param:"id"`
}

type GetBlogResponse struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Writer  string   `json:"writer"`
	View    int      `json:"view"`
	Heart   int      `json:"heart"`
	Tags    []string `json:"tags"`
}

type ServiceGetBlog struct {
	*context.Context
	req GetBlogRequest
}

func (app *ServiceGetBlog) Service() *api_context.CommonResponse {
	tb := database.TBBlog{}
	if err := tb.Get(app.DB, app.req.ID); err != nil {
		app.Log.Errorf("blog get err - blogID : %s", app.req.ID)
		return api_context.FailureJSON(http.StatusInternalServerError, "블로그 상세 정보 불러오기 실패")
	}

	id := app.AuthInfo.GetUserID()
	fmt.Println("session id", id)

	resp := GetBlogResponse{
		ID:      strconv.Itoa(int(tb.ID)),
		Title:   tb.Title,
		Content: tb.Content,
		Writer:  tb.Writer,
		View:    tb.View,
		Heart:   tb.Heart,
		Tags:    make([]string, len(tb.Tags)),
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
