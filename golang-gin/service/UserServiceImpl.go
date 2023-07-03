package service

import (
	"fmt"
	"main/golang-gin/domain"
	"main/golang-gin/exception"
	"main/golang-gin/repository"
	"time"
)

// UserService 인터페이스 구현체
// UserRepository에 의존
type UserServiceImpl struct {
	userRepository repository.UserRepository
}

// 생성자 대체
func NewUserServiceImpl(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository: userRepository}
}

// 회원가입
func (userServiceImpl UserServiceImpl) Join(signUpReqDto *domain.SignUpReqDto) (*domain.UserResDto, error) {
	// 중복검사
	findUser, _ := userServiceImpl.userRepository.FindOneByUserId((signUpReqDto.GetUserId()))
	if findUser != nil {
		return nil, exception.NewUserDuplicatedException()
	}

	// user 객체 생성
	newUser := signUpReqDto.ToUserModel(time.Now())
	fmt.Println(newUser.GetUserId())

	// 저장
	userServiceImpl.userRepository.Save(newUser)

	return newUser.ToResDto(), nil
}

// User 검색 By UserId
func (userServiceImpl UserServiceImpl) FindOneByUserId(userId string) (*domain.UserResDto, error) {
	findUser, err := userServiceImpl.userRepository.FindOneByUserId(userId)
	if err != nil {
		return nil, err
	}

	findUserResDto := findUser.ToResDto()

	return findUserResDto, nil
}

// User 삭제 By UserId
func (userServiceImpl UserServiceImpl) DeleteByUserId(userId string) error {
	// 유저 검색
	if findUser, _ := userServiceImpl.userRepository.FindOneByUserId(userId); findUser == nil {
		return exception.NewUserNotFoundException("id", userId)
	}

	userServiceImpl.userRepository.DeleteByUserId(userId)

	return nil
}
