package repository

import (
	"main/golang-gin/domain"
)

// User 데이터 담당 인터페이스 구현
type UserRepository interface {
	Save(*domain.User) *domain.UserResDto
	FindOneByUserId(userId string) (*domain.User, error)
	DeleteByUserId(userId string) error
}
