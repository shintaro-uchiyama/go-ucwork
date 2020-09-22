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

func (s UserApplicationService) Create(user domain.User) error {

	isExist, err := s.userService.Exists(user)
	if err != nil {
		return fmt.Errorf("[application.Create] check exist error: %w", err)
	}
	if isExist {
		return errors.New("[application.Create] request user already registered")
	}

	saveErr := s.userRepository.Save(&user)
	if saveErr != nil {
		return fmt.Errorf("[application.Create] save user error: %w", saveErr)
	}
	return nil
}

func (s UserApplicationService) List() (domain.Users, error) {
	users, err := s.userRepository.FindAll()
	return users, err
}
