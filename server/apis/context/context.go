package context

import (
	"devLog/common/api_context"
	"devLog/server/config"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"time"
)

const (
	sessionKey    = "auth"
	idSessionKey  = "user_id"
	SessionPath   = "/api/v1"
	SessionSecret = "devLog-session-secret"
	MaxAge        = 3600
)

type Context struct {
	*api_context.Context
	Config     *config.Config
	DB         *gorm.DB
	httpClient *resty.Client
	AuthInfo   AuthHandler
	Emails     Email
}

type Email struct {
	YeongAddress string
	YeongPwd     string
	YujinAddress string
	YujinPwd     string
}

func (c Context) GetAuthHandler() api_context.AuthInterface {
	return c.AuthInfo.GetHandler()
}

type AuthHandler struct {
	UserId    string
	LoginDate time.Time
}

func (a AuthHandler) SetAuthInfoInContext(data interface{}) {
	a = data.(AuthHandler)
}

func (a AuthHandler) ParseAuthorization(c echo.Context) *api_context.CommonResponse {
	sess, err := session.Get(sessionKey, c)
	if err != nil {
		return api_context.FailureJSON(http.StatusUnauthorized, "세션 정보를 찾을 수 없습니다.")
	}

	logrus.Info("sess : ", sess)

	sess.Options = &sessions.Options{
		Path:     SessionPath,
		MaxAge:   MaxAge,
		HttpOnly: true,
	}

	if val, has := sess.Values["auth"]; has == false {
		return api_context.FailureJSON(http.StatusUnauthorized, "세션 정보를 찾을 수 없습니다.")
	} else {
		a.UserId = val.(string)
	}

	if err = sess.Save(c.Request(), c.Response()); err != nil {
		return api_context.FailureJSON(http.StatusInternalServerError, "세션 생성 실패")
	}

	return nil
}

func (a AuthHandler) GetUserID(c echo.Context) string {
	sess, err := session.Get(sessionKey, c)
	if err != nil {
		return ""
	}

	sess.Options = &sessions.Options{
		Path:     SessionPath,
		MaxAge:   MaxAge,
		HttpOnly: true,
	}

	if val, has := sess.Values["auth"]; has == false {
		return ""
	} else {
		a.UserId = val.(string)
	}
	return a.UserId
}

func (a AuthHandler) Get(c echo.Context) (interface{}, error) {
	sess, err := session.Get(sessionKey, c)
	if err != nil {
		return nil, err
	}
	val, has := sess.Values["auth"]
	if has == false {
		logrus.Debugf("middlware session not found user")
		return a, fmt.Errorf("session auth is empty")
	}
	_ = json.Unmarshal(val.([]byte), a)
	return a, nil
}

func (a AuthHandler) Create(c echo.Context, data interface{}) error {
	sess, err := session.Get(sessionKey, c)
	if err != nil {
		return fmt.Errorf("session get failed. : %s", err)
	}
	auth := data.(AuthHandler)
	a.SetSessionAuth(sess, auth)
	logrus.Debugf("create session : %+v", auth)
	return sess.Save(c.Request(), c.Response())
}

func (a AuthHandler) SetSessionAuth(sess *sessions.Session, auth AuthHandler) {
	auth.LoginDate = time.Now()
	b, _ := json.Marshal(&auth)
	sess.Values["auth"] = b
	sess.Options = &sessions.Options{
		Path:     "/api/v1",
		MaxAge:   3600,
		HttpOnly: false,
	}
}

func (a AuthHandler) SetCookie(c echo.Context, userID string) error {
	cookie := new(http.Cookie)
	cookie.Name = "user_id"
	cookie.Value = userID
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return nil
}

func (a AuthHandler) ExpireCookie(c echo.Context) error {
	cookie, err := c.Cookie("user_id")
	if err != nil {
		logrus.WithError(err).Error("no cookie with user_id")
		return err
	}
	cookie.Expires = time.Now()
	c.SetCookie(cookie)
	return nil
}

func (a AuthHandler) GetSessionID(c echo.Context) string {
	sess, err := session.Get(sessionKey, c)
	if err != nil {
		return ""
	}
	sess.Values["auth"] = a
	return a.UserId
}

func (a AuthHandler) Check() error {
	//TODO check
	return nil
}

func (a AuthHandler) GetHandler() api_context.AuthInterface {
	return a
}

func InitContext(e *echo.Echo, logger *logrus.Logger, db *gorm.DB, client *resty.Client, cfg *config.Config) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			ctx := &Context{
				Context:    api_context.NewContext(context, logger, AuthHandler{}),
				Config:     cfg,
				DB:         db,
				httpClient: client,
				Emails: Email{
					YeongAddress: cfg.Email.YeongEmail,
					YeongPwd:     cfg.Email.YeongPwd,
					YujinAddress: cfg.Email.YujinEmail,
					YujinPwd:     cfg.Email.YujinPwd,
				},
			}
			return next(ctx)
		}
	})
	api_context.EnableTransactionLog(e, logger)
}
