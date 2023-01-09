package routers

import (
	"picview/middleware"
	"picview/services"

	"github.com/gin-gonic/gin"
)

func Routers() (router *gin.Engine) {
	router = gin.New()
	router.Use(middleware.DefaultStructuredLogger())
	router.Use(gin.Recovery())

	rgroup := router.Group("/api")

	rgroup.GET("/random", services.Random)
	rgroup.Static("/images", "./images")

	return
}
