package application

import (
	"errors"

	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
)

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

func (s UserApplicationService) Regist(user domain.User) (domain.User, error) {
	if s.userService.Exists(user) {
		return domain.User{}, errors.New("duplicated")
	}

	savedUser, err := s.userRepository.Save(&user)
	return *savedUser, err
}

func (s UserApplicationService) List() (domain.Users, error) {
	users, err := s.userRepository.FindAll()
	return *users, err
}
