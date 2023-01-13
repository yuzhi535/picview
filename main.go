package main

import (
	"os"
	"picview/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("app.env"); err != nil {
		panic("Error loading .env file")
	}
	port := "0.0.0.0:" + os.Getenv("PORT")
	gin_mode := os.Getenv("GIN_MODE")
	gin.SetMode(gin_mode)
	router := routers.Routers()
	router.Run(port)
}
