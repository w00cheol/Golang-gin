package exception

// 잘못된 요청 인자 예외 처리
type IllegalArgumentException struct {
	message string
}

// 생성자
func NewIllegalArgumentException(invalidArg string) *IllegalArgumentException { // todo: parameter 어떻게 하면 좋을까
	message := "argument cannot be " + invalidArg
	return &IllegalArgumentException{message: message}
}

// Error 인터페이스 구현
func (e *IllegalArgumentException) Error() string {
	return e.message
}
