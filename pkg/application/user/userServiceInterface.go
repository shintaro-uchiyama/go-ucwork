package application

import domain "github.com/shintaro-uchiyama/go-ucwork/pkg/domain/user"

type UserServiceInterface interface {
	Exists(user domain.User) bool
}