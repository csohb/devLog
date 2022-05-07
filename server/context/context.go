package context

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type ServiceHandler interface {
	Service() *CommonResponse
	GetRequestData() interface{}
	ApiName() interface{}
	IsRequiredAuth() bool
	GetContext() *Context
	SetLogAPI(name string)
	GetAuthHandler()
}
type Context struct {
	echo.Context
	log      *logrus.Logger
	Log      *logrus.Entry
	HasToken bool
	RespData CommonResponse
	Auth     AuthInterface
}

func (ctx *Context) GetContext() *Context {
	return ctx
}

func (ctx *Context) SetLogAPI(name string) {
	ctx.Log = ctx.Log.WithField("api", name)
}

func Process(service ServiceHandler) error {
	var err error

	c := service.GetContext()

}
