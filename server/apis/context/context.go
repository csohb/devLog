package context

import (
	"devLog/common/api_context"
	"devLog/server/config"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

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
	Grade     int
}

func (a *AuthHandler) ErrorHandler() error {
	//TODO error handler
	return nil
}

func (a *AuthHandler) Check() error {
	//TODO check
	return nil
}

func (a *AuthHandler) Create(c echo.Context, data interface{}) error {
	//TODO implement me
	/*
		jwt token으로 사용자 판단
	*/
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
