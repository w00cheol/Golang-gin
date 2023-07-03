package infrastructure

import (
	"main/golang-gin/controller"
	"main/golang-gin/repository"
	"main/golang-gin/service"

	"github.com/gin-gonic/gin"
)

// 의존성 injecting
// todo: 환경변수 또는 .sh 로 injection configuration 해보자
type Server struct {
}

func (server Server) Inject(router *gin.Engine) *gin.Engine {

	// Databases
	userDb := repository.NewLocalUserDb()

	// Repositorys
	userRepository := repository.NewLocalUserRepository(userDb)

	// Services
	userService := service.NewUserServiceImpl(userRepository)

	// Controllers
	userController := controller.NewUserControllerImpl(userService)

	// Initialize
	userController.Init(router)

	return router
}
