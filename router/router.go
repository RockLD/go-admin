package router

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"go-admin/middleware"
	"go-admin/controllers"
)

func Engine() *gin.Engine{
	router := gin.Default()
	router.GET("/ad/login", func(c *gin.Context) {
		fmt.Println("login")
	})
	router.POST("/ad/do-login", func(c *gin.Context) {
		fmt.Println("do-login")
	})

	//校验
	authorize := router.Group("/admin",middleware.SessionAuth)
	{
		authorize.GET("/index",controllers.Index{}.Index)
		authorize.GET("/first",controllers.Index{}.First)
		authorize.Any("/nodes",controllers.Nodes{}.Index)
		authorize.Any("/admins",controllers.Admins{}.Index)
	}

	return router
}