package application

import (
	"errors"
	"fmt"

	"github.com/shintaro-uchiyama/go-ucwork/pkg/presentation"

	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
)

var _ presentation.UserApplicationServiceInterface = (*UserApplicationService)(nil)

type UserApplicationService struct {
	userRepository domain.UserRepositoryInterface
	userService    UserServiceInterface
}

func NewUserApplicationService(userRepository domain.UserRepositoryInterface, userService UserServiceInterface) *UserApplicationService {
	return &UserApplicationService{
		userRepository: userRepository,
		userService:    userService,
	}
}

func (s UserApplicationService) Create(user domain.User) (*domain.User, error) {

	isExist, err := s.userService.Exists(user)
	if err != nil {
		return nil, fmt.Errorf("[application.Create] check exist error: %w", err)
	}
	if isExist {
		return nil, errors.New("[application.Create] request user already registered")
	}

	s.userRepository.Begin()
	savedUser, saveErr := s.userRepository.Save(&user)
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
