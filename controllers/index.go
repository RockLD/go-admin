package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-admin/models"
)

type IndexController struct {

}

func (index IndexController)Index(c *gin.Context) {
	var np models.Nodes
	nodes,_ := np.GetList()

	//fmt.Println(nodes)
	c.HTML(http.StatusOK,"index/index.html",gin.H{
		"nodes":nodes,
	})
}

func (index IndexController)First(c *gin.Context) {
	c.HTML(http.StatusOK,"index/first.html",gin.H{})
}
