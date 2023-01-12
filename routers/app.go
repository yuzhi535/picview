package routers

import (
	"picview/middleware"
	"picview/services"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers() (router *gin.Engine) {
	router = gin.New()
	router.Use(middleware.DefaultStructuredLogger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*", "https://github.com"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin[:5] == "https" // just an example!
		// },
	}))

	rgroup := router.Group("/api")

	rgroup.GET("/random", services.Random)
	rgroup.Static("/images", "./images")

	return
}
