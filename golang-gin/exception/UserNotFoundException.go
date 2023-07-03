package exception

import (
	"fmt"
)

// 유저 미존재 예외처리
type UserNotFoundException struct {
	keywordType  string
	keywordValue any
	message      string
}

// 생성자
func NewUserNotFoundException(keywordType string, keywordValue any) *UserNotFoundException {
	message := fmt.Sprintf("No Users match this %s", keywordType)
	return &UserNotFoundException{keywordType: keywordType, keywordValue: keywordValue, message: message}
}

// Error 인터페이스 구현
func (e *UserNotFoundException) Error() string {
	return e.message
}
