package application

import (
	"context"

	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
)

type FirebaseAuthInterface interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
}
