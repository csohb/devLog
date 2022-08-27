package email

import (
	"devLog/common/api_context"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
)

type SendEmailRequest struct {
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
}

type ServiceEmailSend struct {
	*context.Context
	req SendEmailRequest
}

func (app *ServiceEmailSend) Service() *api_context.CommonResponse {
	return api_context.SuccessJSON(nil)
}

func (app *ServiceEmailSend) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceEmailSend) ApiName() interface{} {
	return "send-email"
}

func (app ServiceEmailSend) IsRequiredAuth() bool {
	return false
}

func ProcessSendEmail(c echo.Context) error {
	return api_context.Process(&ServiceEmailSend{
		Context: c.(*context.Context),
	})
}
