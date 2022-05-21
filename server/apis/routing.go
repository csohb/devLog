package apis

import (
	"devLog/server/apis/career"
	"devLog/server/apis/mainpage"
	"github.com/labstack/echo/v4"
)

func RoutingForDevLog(grp *echo.Group) {
	// main page
	grp.GET("/introduce/:id", mainpage.ProcessIntroduce)

	// career
	grp.POST("/career/save", career.ProcessSaveCareer)
}
