package controllers

import (
	"github.com/gin-gonic/gin"
	"go-admin/models"
	"net/http"
)

type NodesController struct {

}

func (nodes NodesController) Index(c *gin.Context) {
	nds,err := models.Nodes{}.GetList()

	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"code":2,
			"message":err,
		})
	}

	c.JSON(http.StatusOK,gin.H{
		"code":1,
		"message":"success",
		"data":nds,
	})
}
