package domain

import "time"

// User 생성 시 요청 DTO
type SignUpReqDto struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

// User 모델로 변환
func (signUpReqDto *SignUpReqDto) ToUserModel(createdAt time.Time) *User {
	return &User{UserId: signUpReqDto.UserId, Password: signUpReqDto.Password, CreatedAt: createdAt}
}

// getter 함수
func (signUpReqDto *SignUpReqDto) GetUserId() string {
	return signUpReqDto.UserId
}
