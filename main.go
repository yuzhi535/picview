package main

import (
	"fmt"
	"picview/routers"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	// ...
	router := routers.Routers()
	router.Run(":8080")
}
