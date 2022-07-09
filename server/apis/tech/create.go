package tech

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Tech struct {
	Name       string `json:"name"`
	UserID     string `json:"user_id"`
	Percentage int    `json:"percentage,omitempty"`
}

type CreateTechRequest struct {
	Tech []Tech `json:"tech"`
}

type ServiceCreateTech struct {
	*context.Context
	req CreateTechRequest
}

func (app *ServiceCreateTech) Service() *api_context.CommonResponse {
	tb := database.TBTech{}
	fmt.Println("req :", app.req)
	for _, v := range app.req.Tech {
		tb = database.TBTech{
			Name:       v.Name,
			UserID:     v.UserID,
			Percentage: v.Percentage,
		}
		if err := tb.Create(app.DB); err != nil {
			logrus.WithError(err).Error("create tech failed")
			return api_context.FailureJSON(http.StatusInternalServerError, "db create err")
		}
	}

	return api_context.SuccessJSON(nil)
}

func (app *ServiceCreateTech) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceCreateTech) ApiName() interface{} {
	return "create-tech"
}

func (app ServiceCreateTech) IsRequiredAuth() bool {
	return true
}

func ProcessCreateTech(c echo.Context) error {
	return api_context.Process(&ServiceCreateTech{
		Context: c.(*context.Context),
	})
}
