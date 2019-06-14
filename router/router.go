package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/middleware"
	"go-admin/controllers"
	"net/http"
)

func Engine() *gin.Engine{
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.StaticFile("/favicon.ico","./static/favicon.ico")
	router.Static("/css","./static/css")
	router.Static("/js","./static/js")
	router.Static("/img","./static/img")
	router.Static("/fonts","./static/fonts")
	router.Static("/plugins","./static/plugins")
	router.Static("/layui","./static/layui")

	router.GET("/ad/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login/login.html",gin.H{
			"title":"Go-Admin后台 - 登录",
		})
	})
	router.POST("/ad/do-login", controllers.Login{}.Login)

	//校验
	authorize := router.Group("/admin",middleware.SessionAuth)
	{
		authorize.GET("/index",controllers.IndexController{}.Index)
		authorize.GET("/first",controllers.IndexController{}.First)
		authorize.Any("/nodes",controllers.NodesController{}.Index)
		authorize.Any("/admins",controllers.Admins{}.Index)
	}

	return router
}