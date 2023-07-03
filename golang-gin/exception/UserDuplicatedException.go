package exception

// 유저 중복 예외 처리
type UserDuplicatedException struct {
	message string
}

func NewUserDuplicatedException() *UserDuplicatedException {
	message := "User with this id already exist"
	return &UserDuplicatedException{message: message}
}

// Error 인터페이스 구현
func (e *UserDuplicatedException) Error() string {
	return e.message
}
