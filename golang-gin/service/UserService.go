package service

import "main/golang-gin/domain"

// User 관련 비즈니스 로직 인터페이스
type UserService interface {
	Join(signUpReqDto *domain.SignUpReqDto) (*domain.UserResDto, error)
	FindOneByUserId(userId string) (*domain.UserResDto, error)
	DeleteByUserId(userId string) error
}
