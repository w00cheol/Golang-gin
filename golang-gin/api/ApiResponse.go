package api

import "github.com/gin-gonic/gin"

// Status 상수 정의
const (
	StatusSuccess string = "Success"
	StatusFail    string = "Fail"
)

// JSON 응답 객체
type ApiResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}

// 에러 발생 시 JSON 응답 함수
func ResponseFail(ctx *gin.Context, httpStatus int, err error) {
	ctx.JSON(httpStatus, ApiResponse{
		Status: StatusFail,
		Error:  err.Error(),
	})
}

// 성공 시 JSON 응답 함수
func ResponseSuccess(ctx *gin.Context, httpStatus int, data any) {
	ctx.JSON(httpStatus, ApiResponse{
		Status: StatusSuccess,
		Data:   data,
	})
}
