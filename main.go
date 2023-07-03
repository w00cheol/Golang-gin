/*
https://school.programmers.co.kr/learn/courses/30/lessons/181944?language=go
홀짝 구분하기
*/

package main

import (
	"main/golang-gin/infrastructure"
	"os"

	"github.com/gin-gonic/gin"
)

// 그냥 간단한 분기문
func main() {
	serverPort := os.Args[1]

	router := gin.Default()
	router = infrastructure.Server{}.Inject(router)

	router.Run(":" + serverPort)
}
