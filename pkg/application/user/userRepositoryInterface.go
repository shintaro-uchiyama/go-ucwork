package application

import domain "github.com/shintaro-uchiyama/go-ucwork/pkg/domain/user"

type UserRepositoryInterface interface {
	Save(user *domain.User) (*domain.User, error)
}