package presentation

import (
	"context"

	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
)

type UserApplicationServiceInterface interface {
	Create(ctx context.Context, user domain.User) (*domain.User, error)
	List() ([]domain.User, error)
}
