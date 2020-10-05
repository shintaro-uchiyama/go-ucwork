package application

import (
	"context"
	"errors"
	"fmt"

	"github.com/shintaro-uchiyama/go-ucwork/pkg/presentation"

	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
)

var _ presentation.UserApplicationServiceInterface = (*UserApplicationService)(nil)

type UserApplicationService struct {
	userRepository domain.UserRepositoryInterface
	firebaseAuth   FirebaseAuthInterface
	userService    UserServiceInterface
}

func NewUserApplicationService(
	userRepository domain.UserRepositoryInterface,
	firebaseAuth FirebaseAuthInterface,
	userService UserServiceInterface,
) *UserApplicationService {
	return &UserApplicationService{
		userRepository: userRepository,
		firebaseAuth:   firebaseAuth,
		userService:    userService,
	}
}

func (s UserApplicationService) Create(ctx context.Context, user domain.User) (*domain.User, error) {
	isExist, existErr := s.userService.Exists(user)
	if existErr != nil {
		return nil, fmt.Errorf("[application.Create] check exist existError: %w", existErr)
	}
	if isExist {
		return nil, errors.New("[application.Create] request user already registered")
	}

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
