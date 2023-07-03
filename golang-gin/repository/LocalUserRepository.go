package repository

import (
	"fmt"
	"main/golang-gin/domain"
	"main/golang-gin/exception"
)

// DB 사용하지 않고 구조체 변수 사용
type LocalUserRepository struct {
	population int
	db         map[string]*domain.User
}

// 생성자 함수 대체
func NewLocalUserRepository(db map[string]*domain.User) *LocalUserRepository {
	return &LocalUserRepository{population: 0, db: db}
}

// User 저장함수
func (localUserRepository *LocalUserRepository) Save(user *domain.User) *domain.UserResDto {
	localUserRepository.db[user.GetUserId()] = user
	localUserRepository.population++
	fmt.Println(localUserRepository.population)

	return user.ToResDto()
}

// User 검색 by userId
func (localUserRepository *LocalUserRepository) FindOneByUserId(userId string) (*domain.User, error) {
	if user, ok := localUserRepository.db[userId]; ok {
		return user, nil
	}

	return nil, exception.NewUserNotFoundException("id", userId)
}

// User 삭제 by userId
func (localUserRepository *LocalUserRepository) DeleteByUserId(userId string) error {
	delete(localUserRepository.db, userId)
	localUserRepository.population--
	fmt.Println(localUserRepository.population)

	return nil
}
