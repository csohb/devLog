package career

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DeleteCareerRequest struct {
	ID []int `param:"id"`
}

type ServiceDeleteCareer struct {
	*context.Context
	req DeleteCareerRequest
}

func (app *ServiceDeleteCareer) Service() *api_context.CommonResponse {
	fmt.Println("request data : ", app.req)
	tb := database.TBCareer{}

	if err := tb.DeleteCareer(app.DB, app.req.ID); err != nil {
		app.Log.Errorf("Delete career failed : %+v", err)
		return api_context.FailureJSON(http.StatusInternalServerError, "이력 삭제에 실패하였습니다.")
	}
	return api_context.SuccessJSON(nil)
}

func (app *ServiceDeleteCareer) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceDeleteCareer) ApiName() interface{} {
	return "delete-career"
}

func (app ServiceDeleteCareer) IsRequiredAuth() bool {
	return true
}

func ProcessDeleteCareer(c echo.Context) error {
	return api_context.Process(&ServiceDeleteCareer{Context: c.(*context.Context)})
}
