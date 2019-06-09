package controllers

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"go-admin/common"
	"go-admin/middleware"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (login Login) Login(c *gin.Context) {
	fmt.Println("do-login")
	username := c.PostForm("username")

	if username != "admin" {
		c.JSON(http.StatusOK,gin.H{
			"type":"failed",
			"message":"账号错误",
		})
		return
	}
	password := common.MyMd5(c.PostForm("password"))

	md := "e10adc3949ba59abbe56e057f20f883e"
	if password != md {
		c.JSON(http.StatusOK,gin.H{
			"type":"failed",
			"message":"密码错误",
		})
		return
	}

	session,err := middleware.Store.Get(c.Request,"admin_id")
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"type":"failed",
			"message":"session error",
		})
		return
	}
	session.Values["admin_id"] = 1
	session.Values["username"] = username
	err = session.Save(c.Request,c.Writer)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"type":"failed",
			"message":"session error",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
			"msg":"登录成功！",
			"type":"success",
	})
	return
}
