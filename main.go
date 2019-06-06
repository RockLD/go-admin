package main

import (
	"github.com/gin-gonic/gin"
	"go-admin/controllers"
)
func main() {
	router := gin.Default()
	router.GET("/index/index",controllers.Index{}.Index)

	router.Run(":8090")

}
