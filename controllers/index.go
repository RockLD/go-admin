package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/common"
	"go-admin/middleware"
	"go-admin/models"
	"net/http"
)

type IndexController struct {
}

func (index IndexController) Index(c *gin.Context) {
	session, err := middleware.Store.Get(c.Request, "role_id")
	if err != nil {
		c.String(http.StatusInternalServerError, "session,err")
	}

	role_id := session.Values["role_id"]
	nodes, _ := models.Nodes{}.GetListByRole(role_id.(int))

	nodeTree := common.CreateMenuTree(nodes)
	fmt.Println(nodeTree)
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"nodes": nodeTree,
	})
}

func (index IndexController) First(c *gin.Context) {
	c.HTML(http.StatusOK, "index/first.html", gin.H{})
}
