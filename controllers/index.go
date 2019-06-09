package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Index struct {

}

func (index Index)Index(c *gin.Context) {
	c.HTML(http.StatusOK,"index/index.html",gin.H{
		"nodes":"",
	})
}

func (index Index)First(c *gin.Context) {

}
