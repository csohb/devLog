package login

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/context"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type RequestLogin2 struct {
	UserID   string `json:"id"`
	Password string `json:"password"`
}

type ResponseLogin2 struct {
	Token string `json:"token"`
}

type ServiceLogin2 struct {
	*context.Context
	req RequestLogin
}

type Claim struct {
	UserID string
	jwt.StandardClaims
}

type User2 struct {
	ID  string `json:"id"`
	Pwd string `json:"pwd"`
}

var expirationTime = 30 * time.Minute
var JwtKey = []byte("v_^.^_v")

func (u *User2) GetJwtToken() (string, error) {
	expire := time.Now().Add(expirationTime)
	claims := &Claim{
		UserID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", errors.Wrap(err, "token signed error")
	} else {
		return tokenString, nil
	}
}

func (app *ServiceLogin2) Service() *api_context.CommonResponse {
	// id find
	tb := database.TBUser{}
	if err := tb.CheckID(app.DB, app.req.UserID); err != nil {
		app.Log.WithError(err).Error("not exist id")
		return api_context.FailureJSON(http.StatusBadRequest, "존재하지 않는 아이디입니다.")
	}
	// pwd check
	logrus.Infof("req.ID : %+v", app.req.UserID)
	logrus.Infof("req.password : %+v", app.req.Password)

	if err := tb.CheckPwd(app.DB, app.req.UserID, app.req.Password); err != nil {
		app.Log.WithError(err).Error("not correct pwd")
		return api_context.FailureJSON(http.StatusBadRequest, "비밀번호가 일치하지 않습니다.")
	}

	resp := ResponseLogin2{}

	return api_context.SuccessJSON(&resp)
}

func (app *ServiceLogin2) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceLogin2) ApiName() interface{} {
	return "devLog-login"
}

func (app ServiceLogin2) IsRequiredAuth() bool {
	return false
}

func ProcessLogin2(c echo.Context) error {
	return api_context.Process(&ServiceLogin{
		Context: c.(*context.Context),
	})
}
