package introduce

import (
	"devLog/common/api_context"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
)

type DeleteIntroduceRequest struct {
	ID int `parma:"id"`
}

type ServiceDeleteIntroduce struct {
	*context.Context
	req DeleteIntroduceRequest
}

func (app *ServiceDeleteIntroduce) Service() *api_context.CommonResponse {
	//TODO implement me
	panic("implement me")
}

func (app *ServiceDeleteIntroduce) GetRequestData() interface{} {
	//TODO implement me
	panic("implement me")
}

func (app ServiceDeleteIntroduce) ApiName() interface{} {
	//TODO implement me
	panic("implement me")
}

func (app ServiceDeleteIntroduce) IsRequiredAuth() bool {
	//TODO implement me
	panic("implement me")
}

func ProcessDeleteIntroduce(c echo.Context) error {
	return api_context.Process(&ServiceDeleteIntroduce{
		Context: c.(*context.Context),
	})
}
