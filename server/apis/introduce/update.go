package introduce

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UpdateProfi struct {
	NickName string `json:"nick_name"`
	Img      string `json:"img"`
}

type UpdateCareer struct {
	Company   string `json:"company"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	JobTitle  string `json:"job_title"`
	JobDetail string `json:"job_detail"`
}

type UpdateProject struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	IsPersonal  bool   `json:"is_personal"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
}

type UpdateIntroduceRequest struct {
	UserID  string          `json:"user_id"`
	Profile UpdateProfi     `json:"profile"`
	Career  []UpdateCareer  `json:"career"`
	Project []UpdateProject `json:"project"`
}

type ServiceUpdateIntroduce struct {
	*context.Context
	req UpdateIntroduceRequest
}

func (app *ServiceUpdateIntroduce) Service() *api_context.CommonResponse {
	tb := database.TBUser{}
	if err := tb.UpdateIntroduce(app.DB, app.req.UserID); err != nil {
		app.Log.Error("update introduce failed.")
		return api_context.FailureJSON(http.StatusInternalServerError, "update db err")
	}

	return api_context.SuccessJSON(nil)
}

func (app *ServiceUpdateIntroduce) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceUpdateIntroduce) ApiName() interface{} {
	return "introduce-update"
}

func (app ServiceUpdateIntroduce) IsRequiredAuth() bool {
	return true
}

func ProcessUpdateIntroduce(c echo.Context) error {
	return api_context.Process(&ServiceUpdateIntroduce{
		Context: c.(*context.Context),
	})
}
