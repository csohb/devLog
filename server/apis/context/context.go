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
	"time"
)

const sessionKey = "devLogSessionKey"

type Context struct {
	*api_context.Context
	Config     *config.Config
	DB         *gorm.DB
	httpClient *resty.Client
	AuthInfo   AuthHandler
}

func (c Context) GetAuthHandler() api_context.AuthInterface {
	return c.AuthInfo.GetHandler()
}

type AuthHandler struct {
	UserId    string
	LoginDate time.Time
}

func (a *AuthHandler) ErrorHandler() error {
	//TODO error handler
	return nil
}

func (a AuthHandler) GetUserID() string {
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
	c.Set("auth", &auth)
	logrus.Debugf("create session : %+v", auth)
	return sess.Save(c.Request(), c.Response())
}

func (a AuthHandler) SetSessionAuth(sess *sessions.Session, auth AuthHandler) {
	auth.LoginDate = time.Now()
	b, _ := json.Marshal(&auth)
	sess.Values["auth"] = b
	sess.Options = &sessions.Options{
		Path:     "/api",
		MaxAge:   3600,
		HttpOnly: false,
	}
}

func (a AuthHandler) GetSessionID(sess *sessions.Session) string {
	sess.Values["auth"] = a
	return a.UserId
}

func (a *AuthHandler) Check() error {
	//TODO check
	return nil
}

func (a *AuthHandler) GetHandler() api_context.AuthInterface {
	return a
}

/*
func InitContext(e *echo.Echo, logger *logrus.Logger, db *gorm.DB, client *resty.Client, cfg *config.Config) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &Context{
				Context:    common.NewContext(c, logger),
				DB:         db,
				httpClient: client,
				AuthInfo:   AuthHandler{},
				Config:     cfg,
			}
			return next(ctx)
		}
	})
	common.EnableTransactionLog(e, logger)
}*/

func InitContext(e *echo.Echo, logger *logrus.Logger, db *gorm.DB, client *resty.Client, cfg *config.Config) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			ctx := &Context{
				Context:    api_context.NewContext(context, logger),
				DB:         db,
				httpClient: client,
				Config:     cfg,
			}
			return next(ctx)
		}
	})
	api_context.EnableTransactionLog(e, logger)
}
