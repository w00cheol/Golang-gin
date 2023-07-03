package controller

import (
	"fmt"
	"main/golang-gin/api"
	"main/golang-gin/domain"
	"main/golang-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller 인터페이스 구현체
type UserControllerImpl struct {
	userService service.UserService
}

// 생성자 구현
func NewUserControllerImpl(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{userService: userService}
}

// 라우터 REST API URL + Handler Func 적용
func (userControllerImpl *UserControllerImpl) Init(router *gin.Engine) *gin.Engine {
	// todo: 이 부분 더 나눌 수 있을지 고민
	router.GET("/users/:userId", userControllerImpl.GetOneByUserId)
	router.POST("/users", userControllerImpl.Create)
	router.DELETE("/users/:userId", userControllerImpl.Delete)

	return router
}

// User 생성 Handler
func (userController *UserControllerImpl) Create(ctx *gin.Context) {
	userDto := &domain.SignUpReqDto{}
	if err := ctx.Bind(userDto); err != nil {
		api.ResponseFail(ctx, http.StatusBadRequest, err)
		return
	}

	createUser, err := userController.userService.Join(userDto)
	if err != nil {
		api.ResponseFail(ctx, http.StatusInternalServerError, err)
		return
	}

	api.ResponseSuccess(ctx, http.StatusCreated, createUser)
	return
}

// User 삭제 Handler
func (userController *UserControllerImpl) Delete(ctx *gin.Context) {
	userId := ctx.Param("userId")

	if err := userController.userService.DeleteByUserId(userId); err != nil {
		api.ResponseFail(ctx, http.StatusInternalServerError, err)
		return
	}

	api.ResponseSuccess(ctx, http.StatusNoContent, nil)
	return
}

// User 검색 By Id
func (userController *UserControllerImpl) GetOneByUserId(ctx *gin.Context) {
	userId := ctx.Param("userId")

	findUser, err := userController.userService.FindOneByUserId(userId)
	if err != nil {
		api.ResponseFail(ctx, http.StatusInternalServerError, err)
		return
	}
	fmt.Printf("%s success found\n", userId)

	api.ResponseSuccess(ctx, http.StatusOK, findUser)
	return
}
