package repository

import "main/golang-gin/domain"

// 런타임용 User 적재 자료형
func NewLocalUserDb() map[string]*domain.User {
	localUserDb := make(map[string]*domain.User)
	return localUserDb
}
