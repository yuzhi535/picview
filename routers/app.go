package routers

import (
	"net/http"
	"picview/services"

	"github.com/gin-gonic/gin"
)

func Routers() (router *gin.Engine) {

	router = gin.Default()

	rgroup := router.Group("/api")

	rgroup.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	rgroup.GET("/random", services.Random)
	rgroup.Static("/images", "./images")
	return
}
