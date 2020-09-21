package presentation

import (
	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
)

type UserApplicationServiceInterface interface {
	Regist(user domain.User) (domain.User, error)
}

