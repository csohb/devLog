package apis

import (
	"devLog/server/apis/mainpage"
	"github.com/labstack/echo/v4"
)

func RoutingForDevLog(grp *echo.Group) {
	grp.GET("/introduce/:id", mainpage.ProcessIntroduce)
}
