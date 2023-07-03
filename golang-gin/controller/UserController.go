package controller

import (
	"github.com/gin-gonic/gin"
)

// Controller 인터페이스
type UserController interface {
	Init()
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetOne(ctx *gin.Context)
}
