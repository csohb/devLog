package blog

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type FindBlogRequest struct {
	Page  int `query:"page"`
	Count int `query:"count"`
}

type Blog struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Writer  string   `json:"writer"`
	View    int      `json:"view"`
	Heart   int      `json:"heart"`
	Tags    []string `json:"tags"`
}

type FindBlogResponse struct {
	List  []Blog `json:"list"`
	Total int    `json:"total"`
}

type ServiceBlogFind struct {
	*context.Context
	req FindBlogRequest
}

func (app *ServiceBlogFind) Service() *api_context.CommonResponse {
	tb := database.TBBlog{}
	list, total, err := tb.Find(app.DB, app.req.Page-1, app.req.Count)
	if err != nil {
		app.Log.Errorf("find db err")
		return api_context.FailureJSON(http.StatusInternalServerError, "db 조회 에러")
	}
	resp := FindBlogResponse{}
	resp.Total = int(total)
	resp.List = make([]Blog, len(list))

	for i, v := range list {
		tags := make([]string, len(v.Tags))
		for j, l := range v.Tags {
			tags[j] = l.Tag
		}
		resp.List[i] = Blog{
			ID:      strconv.Itoa(int(v.ID)),
			Title:   v.Title,
			Content: v.Content,
			Writer:  v.Writer,
			View:    v.View,
			Heart:   v.Heart,
			Tags:    tags,
		}
	}
	return api_context.SuccessJSON(&resp)
}

func (app *ServiceBlogFind) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceBlogFind) ApiName() interface{} {
	return "blog-find"
}

func (app ServiceBlogFind) IsRequiredAuth() bool {
	return false
}

func ProcessFindBlog(c echo.Context) error {
	return api_context.Process(&ServiceBlogFind{
		Context: c.(*context.Context),
	})
}