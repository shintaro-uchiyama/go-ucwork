package model

import (
	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
)

func (u NewUser) ToDomainUser() *domain.User {
	user := domain.User{}
	user.SetEmail(u.Email)
	user.SetPassword(u.Password)
	return &user
}

func MapToModelUser(u *domain.User) *User {
	return &User{
		UUID:  u.UUID(),
		Email: u.Email(),
	}
}

func MapToModelUsers(u []domain.User) []*User {
	var users []*User
	for _, user := range u {
		users = append(users, MapToModelUser(&user))
	}
	return users
}
