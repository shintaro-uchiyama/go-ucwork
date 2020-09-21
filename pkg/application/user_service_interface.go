package application

import (
	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
)

type UserServiceInterface interface {
	Exists(user domain.User) bool
}