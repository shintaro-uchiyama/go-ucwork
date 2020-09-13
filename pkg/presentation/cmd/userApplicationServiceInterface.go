package presentation

import domain "github.com/shintaro-uchiyama/go-ucwork/pkg/domain/user"

type UserApplicationServiceInterface interface {
	Regist(user domain.User) (domain.User, error)
}

