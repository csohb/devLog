package main

import (
	"bytes"
	"devLog/database/conn"
	"devLog/server/apis"
	"devLog/server/apis/context"
	"devLog/server/config"
	"flag"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"html/template"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const CookieKey = "devLog-cookie"

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
	//db, err := conn.ConnectForTest()
	if err != nil {
		logrus.WithError(err).Error("db connect failed. ")
		os.Exit(-1)
	}

	if err = makeIndexHTML(cfg.Web.PublicPath, "index.html"); err != nil {
		logrus.WithError(err).Error("make index html failed.")
		panic(err)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(CookieKey))))

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
	/*e.GET("/", func(c echo.Context) error {
		file := cfg.Web.PublicPath + "/index.html"
		logrus.Debugf("file : %+v", file)
		return c.File(file)
	})*/

	e.Static("/", cfg.Web.PublicPath)
	e.Static("/static/devLog", cfg.Web.PublicPath)

	apis.RoutingForDevLog(grp)
	request := resty.New()
	context.InitContext(e, logrus.StandardLogger(), db, request, cfg)

	e.Start("localhost:8081")

}

type IndexTemplate struct {
	BuildTS int64
}

func (tmpl IndexTemplate) ParseWrite(filename string) ([]byte, error) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, tmpl); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var indexHTMLTmpl = IndexTemplate{
	BuildTS: time.Now().Unix(),
}

func makeIndexHTML(path, htmlFile string) error {
	indexHTMLTmpl = IndexTemplate{BuildTS: time.Now().Unix()}
	if data, err := indexHTMLTmpl.ParseWrite(path + "/template/" + htmlFile); err != nil {
		logrus.Panic("index html template parse failed : ", err)
		return err
	} else {
		if err = ioutil.WriteFile(path+"/"+htmlFile, data, 0644); err != nil {
			logrus.Panicf("index html template write failed : %s\n%s", err, string(data))
			return err
		}
		return nil
	}
}
