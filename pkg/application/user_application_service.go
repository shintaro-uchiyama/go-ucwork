package application

import (
	"context"
	"fmt"

	"github.com/shintaro-uchiyama/go-ucwork/pkg/presentation"

	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
)

var _ presentation.UserApplicationServiceInterface = (*UserApplicationService)(nil)

type UserApplicationService struct {
	userRepository domain.UserRepositoryInterface
	firebaseAuth   domain.FirebaseAuthInterface
}

func NewUserApplicationService(
	userRepository domain.UserRepositoryInterface,
	firebaseAuth domain.FirebaseAuthInterface,
) *UserApplicationService {
	return &UserApplicationService{
		userRepository: userRepository,
		firebaseAuth:   firebaseAuth,
	}
}

func (s UserApplicationService) Create(ctx context.Context, user domain.User) (*domain.User, error) {
	firebaseUser, firebaseErr := s.firebaseAuth.CreateUser(ctx, &user)
	if firebaseErr != nil {
		return nil, fmt.Errorf("[application.Create] firebase create user error: %w", firebaseErr)
	}

	s.userRepository.Begin()
	savedUser, saveErr := s.userRepository.Save(firebaseUser)
	if saveErr != nil {
		s.userRepository.Rollback()
		return nil, fmt.Errorf("[application.Create] save user error: %w", saveErr)
	}
	s.userRepository.Commit()
	return savedUser, nil
}

func (s UserApplicationService) List() ([]domain.User, error) {
	users, err := s.userRepository.FindAll()
	return users, err
}
