package domain

import "time"

// User 응답 객체
type UserResDto struct {
	UserId    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}
