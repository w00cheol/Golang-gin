package main

import (
	"main/golang-gin/infrastructure"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	serverPort := os.Args[1]

	router := gin.Default()
	router = infrastructure.Server{}.Inject(router)

	router.Run(":" + serverPort)
}
