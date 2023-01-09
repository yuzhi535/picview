package routers

import (
	"picview/services"

	"github.com/gin-gonic/gin"
)

func Routers() (router *gin.Engine) {
	router = gin.Default()

	rgroup := router.Group("/api")

	rgroup.GET("/random", services.Random)
	rgroup.Static("/images", "./images")
	return
}
