package blog

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type DeleteBlogRequest struct {
	ID string `param:"id"`
}

type ServiceDeleteBlog struct {
	*context.Context
	req DeleteBlogRequest
}

func (app *ServiceDeleteBlog) Service() *api_context.CommonResponse {
	id, err := strconv.Atoi(app.req.ID)
	if err != nil {
		return api_context.FailureJSON(http.StatusBadRequest, "id convert to int error")
	}

	err = app.DB.Delete(&database.TBBlog{}, id).Error
	if err != nil {
		app.Log.Errorf("delete blog err : %+v", err)
		return api_context.FailureJSON(http.StatusInternalServerError, "db delete err")
	}

	return api_context.SuccessJSON(nil)
}

func (app *ServiceDeleteBlog) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceDeleteBlog) ApiName() interface{} {
	return "blog-delete"
}

func (app ServiceDeleteBlog) IsRequiredAuth() bool {
	return true
}

func ProcessDeleteBlog(c echo.Context) error {
	return api_context.Process(&ServiceDeleteBlog{
		Context: c.(*context.Context),
	})
}
