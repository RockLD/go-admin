package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-admin/models"
	"go-admin/common"
	"fmt"
)

type IndexController struct {

}

func (index IndexController)Index(c *gin.Context) {
	nodes,_ := models.Nodes{}.GetListByRole()

	nodeTree := common.CreateMenuTree(nodes)
	fmt.Println(nodeTree)
	c.HTML(http.StatusOK,"index/index.html",gin.H{
		"nodes":nodeTree,
	})
}

func (index IndexController)First(c *gin.Context) {
	c.HTML(http.StatusOK,"index/first.html",gin.H{})
}
