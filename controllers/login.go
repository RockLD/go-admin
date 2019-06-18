package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-admin/common"
	"go-admin/middleware"
	"go-admin/models"
)

type LoginController struct {}

//do-login
func (login LoginController) Login(c *gin.Context) {
	username := c.PostForm("username")
	userInfo,err := models.Admins{}.GetInfoByUser(username)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"type":"failed",
			"message":err.Error(),
		})
	}

	if userInfo == nil {
		c.JSON(http.StatusOK,gin.H{
			"type":"failed",
			"message":"账号不存在",
		})
	}

	if userInfo.Status != 1 {
		c.JSON(http.StatusOK,gin.H{
			"type":"failed",
			"message":"账号状态异常，请联系管理员",
		})
	}

	password := common.MyMd5(c.PostForm("password"))

	if password != userInfo.Password {
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
			"message":"session error:" + err.Error(),
		})
		return
	}
	session.Values["admin_id"] = userInfo.AdminId
	session.Values["username"] = username
	session.Values["role_id"] = userInfo.RoleId

	err = session.Save(c.Request,c.Writer)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"type":"failed",
			"message":"session error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
			"msg":"登录成功！",
			"type":"success",
	})
	return
}
