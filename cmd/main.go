package main

import (
	"devLog/database/conn"
	"devLog/server/apis"
	"devLog/server/apis/context"
	"devLog/server/config"
	"flag"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

const Env = ".env"
const Config = "DEV_LOG_CONFIG"

var (
	configFile = flag.String("c", "cms.yaml", "The path to the config file.")
)

func ProcessIgnoreSignal() {
	signal.Ignore(syscall.SIGHUP)
}

func main() {

	// config 다시 설정
	flag.Parse()
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		logrus.WithError(err).Error("Load Config err.")
		os.Exit(-1)
	}

	e := echo.New()
	defer e.Close()

	// db connect
	db, err := conn.ConnectForYJ()
	if err != nil {
		logrus.WithError(err).Error("db connect failed. ")
		os.Exit(-1)
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:          nil,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "DELETE", "OPTION", "PUT", "HEAD"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    nil,
		MaxAge:           0,
	}))

	grp := e.Group("/api/v1")
	apis.RoutingForDevLog(grp)
	request := resty.New()
	context.InitContext(e, logrus.StandardLogger(), db, request, cfg)

	e.Start("localhost:8080")

}
