package email

import (
	"devLog/common/api_context"
	"devLog/server/apis/context"
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/mail"
	"net/smtp"
)

type SendEmailRequest struct {
	Receiver string `json:"receiver"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Number   string `json:"number,omitempty"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type ServiceEmailSend struct {
	*context.Context
	req SendEmailRequest
}

func SendEmail(receiver, sender, pwd string) error {
	auth := smtp.PlainAuth("", sender, pwd, "smtp.gmail.com")
	from := sender
	to := []string{receiver}

	headerSubject := "Subject: 문의 남겨주셔서 감사합니다."
	headerBlank := "\r\n"
	body := "빠른 시일 안에 문의 내용에 답변 드릴 수 있도록 하겠습니다."
	msg := []byte(headerSubject + headerBlank + body)

	fmt.Println("msg : ", string(msg))

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		return err
	}
	return nil
}

func ReceiveMail(name, email, number, title, content string, receiver, pwd string) error {
	auth := smtp.PlainAuth("", receiver, pwd, "smtp.gmail.com")
	from := mail.Address{
		Name:    "yujin",
		Address: receiver,
	}

	to := mail.Address{
		Name:    "yujin",
		Address: receiver,
	}

	body := fmt.Sprintf("문의자 이름:%s 문의자 메일:%s 문의자 번호:%s  문의내용:%s", name, email, number, content)

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = fmt.Sprintf("문의 메일이 도착하였습니다. 제목[%s]", title)

	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))
	err := smtp.SendMail("smtp.gmail.com:587", auth, from.Address, []string{to.Address}, []byte(message))
	if err != nil {
		return err
	}
	return nil
}

func (app *ServiceEmailSend) Service() *api_context.CommonResponse {
	var email string
	var pwd string
	if app.req.Receiver == "yeong" {
		email = app.Emails.YeongAddress
		pwd = app.Emails.YeongPwd
	} else if app.req.Receiver == "yujin" {
		email = app.Emails.YujinAddress
		pwd = app.Emails.YujinPwd
	} else {
		app.Log.Error("not valid sender id")
		return api_context.FailureJSON(http.StatusBadRequest, "존재하는 아이디가 아닙니다.")
	}

	fmt.Println("email : ", email)

	if err := ReceiveMail(app.req.Name, app.req.Email, app.req.Number, app.req.Title, app.req.Content, email, pwd); err != nil {
		app.Log.WithError(err).Error("receive email err")
		return api_context.FailureJSON(http.StatusInternalServerError, "이메일 수신 실패.")
	}

	if err := SendEmail(app.req.Email, email, pwd); err != nil {
		app.Log.WithError(err).Error("send email err")
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
