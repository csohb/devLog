package login

import (
	"devLog/common/api_context"
	"devLog/database"
	"devLog/server/apis/auth"
	"devLog/server/apis/context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type RequestLogin struct {
	UserID   string `json:"id"`
	Password string `json:"password"`
}

type ResponseLogin struct {
	UserID string `json:"user_id"`
}

type ServiceLogin struct {
	*context.Context
	req RequestLogin
}

type Claim struct {
	UserID string
	jwt.StandardClaims
}

type User struct {
	ID  string `json:"id"`
	Pwd string `json:"pwd"`
}

var expirationTime = 30 * time.Minute
var JwtKey = []byte("v_^.^_v")

func (u *User) GetJwtToken() (string, error) {
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

func (app *ServiceLogin) Service() *api_context.CommonResponse {

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

	/*user := User{
		ID:  app.req.UserID,
		Pwd: app.req.Password,
	}

	// 이 토큰 값을 리턴해서 set-cookie에 넣고 response
	jwtToken, err := user.GetJwtToken()
	if err != nil {
		logrus.WithError(err).Error("get jwt token err")
		return api_context.InternalServerErrorJSON("인증 토큰 생성에 실패했습니다.")
	}

	fmt.Println(jwtToken)
	*/
	// auth info create
	authInfo := context.AuthHandler{
		UserId:    app.req.UserID,
		LoginDate: time.Now(),
	}
	app.AuthInfo.Create(app.Context, authInfo)
	fmt.Println("authInfo : ", authInfo)

	sessionAuthInfo, err := auth.CreateSession(app.Context, app.req.UserID)
	if err != nil {
		app.Log.WithError(err).Error("session create failed.")
		return api_context.FailureJSON(http.StatusInternalServerError, "세션 생성 실패")
	}

	resp := ResponseLogin{}
	resp.UserID = sessionAuthInfo.UserID

	fmt.Println("login id : ", sessionAuthInfo.UserID)
	return api_context.SuccessJSON(&resp)
}

func (app *ServiceLogin) GetRequestData() interface{} {
	return &app.req
}

func (app ServiceLogin) ApiName() interface{} {
	return "devLog-login"
}

func (app ServiceLogin) IsRequiredAuth() bool {
	return false
}

func ProcessLogin(c echo.Context) error {
	return api_context.Process(&ServiceLogin{
		Context: c.(*context.Context),
	})
}
