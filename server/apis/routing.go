package apis

import (
	"devLog/server/apis/blog"
	"devLog/server/apis/career"
	"devLog/server/apis/introduce"
	"devLog/server/apis/login"
	"devLog/server/apis/mainpage"
	"devLog/server/apis/project"
	"devLog/server/apis/tech"
	"github.com/labstack/echo/v4"
)

func RoutingForDevLog(grp *echo.Group) {
	// main page
	grp.GET("/main/introduce", mainpage.ProcessProfileList)
	grp.GET("/main/blog", mainpage.ProcessNewestBlogList)

	// login
	grp.POST("/login", login.ProcessLogin)
	grp.POST("/logout", login.ProcessLogout)

	// career
	grp.POST("/career/save", career.ProcessSaveCareer)
	grp.DELETE("/career/delete/:id", career.ProcessDeleteCareer)
	grp.PUT("/career/update", career.ProcessUpdateCareer)

	// blog
	grp.GET("/blog/get/:id", blog.ProcessGetBlog)
	grp.GET("/blog/list", blog.ProcessFindBlog)
	grp.DELETE("/blog/delete/:id", blog.ProcessDeleteBlog)
	grp.POST("/blog/save", blog.ProcessSaveBlog)
	grp.PUT("/blog/heart", blog.ProcessUpdateHeart)
	grp.PUT("/blog/view", blog.ProcessUpdateView)
	grp.PUT("/blog/update", blog.ProcessUpdateBlog)
	grp.GET("/blog/tags", blog.ProcessSearchWithTag)

	// introduce
	grp.GET("/introduce/:id", introduce.ProcessGetIntroduce)
	grp.PUT("/introduce/update", introduce.ProcessUpdateIntroduce)

	// tech
	grp.POST("/introduce/tech/create", tech.ProcessCreateTech)

	// project
	grp.POST("/introduce/project/create", project.ProcessCreateProject)
	grp.DELETE("/introduce/project/delete/:id", project.ProcessDeleteProject)
	grp.PUT("/introduce/project/update", project.ProcessUpdateProject)

}
