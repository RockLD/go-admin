package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	Username string
	Password string
}

func (login Login) Login(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"msg":"success",
		"type":"success",
	})
}
