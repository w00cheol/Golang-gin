package domain

import "time"

// 완전한 User 객체
type User struct {
	UserId    string    `json:"userId"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

// UserResDto로 변환
func (user *User) ToResDto() *UserResDto {
	return &UserResDto{UserId: user.UserId, CreatedAt: user.CreatedAt}
}

// getter 함수
func (user *User) GetUserId() string {
	return user.UserId
}
