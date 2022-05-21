package api_context

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"time"
)

func EnableTransactionLog(e *echo.Echo, logger *logrus.Logger) {
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody []byte, respBody []byte) {
		var reqTime time.Time
		if t := c.Get("requestTime"); t == nil {
			reqTime = time.Now()
		} else {
			reqTime = t.(time.Time)
		}

		req := c.Request()
		c.Request()

		logger.Infof("\n\t[Request  :client->server] ip:%s method:%s path:%s body:\n%s"+
			"\n\t[Response :server->client] status:%d latency:%s body:\n%s",
			c.RealIP(), req.Method, req.RequestURI, string(reqBody),
			c.Response().Status, time.Now().Sub(reqTime).String(), string(respBody),
		)

	}))
}

func NewContext(c echo.Context, logger *logrus.Logger) *Context {
	c.Set("requestTime", time.Now())
	return &Context{
		Context:  c,
		log:      logger,
		Log:      logger.WithField("uri", c.Request().RequestURI),
		HasToken: false,
		RespData: CommonResponse{},
		Auth:     nil,
	}
}
