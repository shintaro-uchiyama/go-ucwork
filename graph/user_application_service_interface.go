package graph

import (
	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
)

type UserApplicationServiceInterface interface {
	Create(user domain.User) (*domain.User, error)
	List() ([]domain.User, error)
}
