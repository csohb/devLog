package email

import (
	"devLog/common/api_context"
	"devLog/server/apis/context"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/smtp"
)

type SendEmailRequest struct {
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
}

type ServiceEmailSend struct {
	*context.Context
	req SendEmailRequest
}

func SendEmail(receiver, sender, pwd string) error {
	auth := smtp.PlainAuth("", sender, pwd, "smtp.gmail.com")
	from := sender
	to := []string{receiver}

	headerSubject := "Subject: [제목] 메일 테스트 발송\r\n"
	headerBlank := "\r\n"
	body := "[본문] 메일 테스트 발송"
	msg := []byte(headerSubject + headerBlank + body)

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		return err
	}
	return nil
}

func (app *ServiceEmailSend) Service() *api_context.CommonResponse {
	var email string
	var pwd string
	if app.req.Sender == "yeong" {
		email = app.Emails.YeongAddress
		pwd = app.Emails.YeongPwd
	} else if app.req.Sender == "yujin" {
		email = app.Emails.YujinAddress
		pwd = app.Emails.YujinPwd
	} else {
		app.Log.Error("not valid sender id")
		return api_context.FailureJSON(http.StatusBadRequest, "존재하는 아이디가 아닙니다.")
	}

	if err := SendEmail(app.req.Receiver, email, pwd); err != nil {
		app.Log.Error("send email err")
		return api_context.FailureJSON(http.StatusInternalServerError, "이메일 전송 실패.")
	}

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
