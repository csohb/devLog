package apis

import (
	"devLog/server/apis/blog"
	"devLog/server/apis/career"
	"devLog/server/apis/mainpage"
	"github.com/labstack/echo/v4"
)

func RoutingForDevLog(grp *echo.Group) {
	// main page
	grp.GET("/introduce/:id", mainpage.ProcessIntroduce)
	grp.GET("/main/introduce", mainpage.ProcessProfileList)

	// career
	grp.POST("/career/save", career.ProcessSaveCareer)
	grp.DELETE("/career/delete", career.ProcessDeleteCareer)
	grp.PUT("/career/update", career.ProcessUpdateCareer)

	// blog
	grp.GET("/blog/get/:id", blog.ProcessGetBlog)
	grp.GET("/blog/list", blog.ProcessFindBlog)
	grp.DELETE("/blog/delete/:id", blog.ProcessDeleteBlog)
	grp.POST("/blog/save", blog.ProcessSaveBlog)
	grp.PUT("/blog/heart", blog.ProcessUpdateHeart)
	grp.PUT("/blog/view", blog.ProcessUpdateView)
	grp.POST("/blog/update", blog.ProcessUpdateBlog)
}
