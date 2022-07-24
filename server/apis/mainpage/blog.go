package mainpage

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Blog struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Date    string   `json:"date"`
	Writer  string   `json:"writer"`
	Tags    []string `json:"tags"`
}

type GetNewestBlogListResponse struct {
	List []Blog `json:"list"`
}

type ServiceGetNewestBlogList struct {
	*context.Context
}

func (app *ServiceGetNewestBlogList) Service() *api_context.CommonResponse {
	// 최신 5개정도?
	tb := database.TBBlog{}
	list, err := tb.BlogForMain(app.DB)
	if err != nil {
		logrus.WithError(err).Error("get list DB failed.")
		return api_context.FailureJSON(http.StatusInternalServerError, "db error")
	}

	resp := GetNewestBlogListResponse{}
	resp.List = make([]Blog, 5)
	for i, v := range list {
		resp.List[i] = Blog{
			ID:      strconv.Itoa(int(v.ID)),
			Title:   v.Title,
			Content: v.Content,
			Date:    v.Model.CreatedAt.Format("2006-01-02"),
			Writer:  v.Writer,
			Tags:    make([]string, len(v.Tags)),
		}
		for j, k := range v.Tags {
			resp.List[i].Tags[j] = k.Tag
		}
	}

	logrus.Debugln("resp.List : ", resp.List)
	return api_context.SuccessJSON(&resp)
}

func (app *ServiceGetNewestBlogList) GetRequestData() interface{} {
	return nil
}

func (app ServiceGetNewestBlogList) ApiName() interface{} {
	return "mainpage-blog"
}

func (app ServiceGetNewestBlogList) IsRequiredAuth() bool {
	return false
}

func ProcessNewestBlogList(c echo.Context) error {
	return api_context.Process(&ServiceGetNewestBlogList{
		Context: c.(*context.Context),
	})
}
