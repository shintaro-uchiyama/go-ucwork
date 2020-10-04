package model

import (
	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
)

func (u NewUser) ToDomainUser() *domain.User {
	user := domain.User{}
	user.SetEmail(u.Email)
	user.SetPassword(u.Password)
	user.SetFirstName(u.FirstName)
	user.SetLastName(u.LastName)
	return &user
}

func MapToModelUser(u *domain.User) *User {
	return &User{
		UUID:      u.UUID(),
		Email:     u.Email(),
		FirstName: u.FirstName(),
		LastName:  u.LastName(),
	}
}

func MapToModelUsers(u []domain.User) []*User {
	var users []*User
	for _, user := range u {
		users = append(users, MapToModelUser(&user))
	}
	return users
}
