package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/controllers"
	"go-admin/middleware"
	"net/http"
)

func Engine() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.Static("/css", "./static/css")
	router.Static("/js", "./static/js")
	router.Static("/img", "./static/img")
	router.Static("/fonts", "./static/fonts")
	router.Static("/plugins", "./static/plugins")
	router.Static("/layui", "./static/layui")

	router.GET("/ad/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login/login.html", gin.H{
			"title": "Go-Admin后台 - 登录",
		})
	})
	router.POST("/ad/do-login", controllers.LoginController{}.Login)

	//校验
	sessionAuth := router.Group("/admin", middleware.SessionAuth)
	{
		sessionAuth.GET("/index", controllers.IndexController{}.Index)
		sessionAuth.GET("/first", controllers.IndexController{}.First)
		sessionAuth.Any("/nodes", controllers.NodesController{}.Index)
		sessionAuth.Any("/admins", controllers.Admins{}.Index)
	}

	return router
}
