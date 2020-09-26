package application

import (
	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
)

var _ UserServiceInterface = (*domain.UserService)(nil)

type UserServiceInterface interface {
	Exists(user domain.User) (bool, error)
}
