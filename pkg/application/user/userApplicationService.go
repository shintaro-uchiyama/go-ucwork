package application

import domain "github.com/shintaro-uchiyama/go-ucwork/pkg/domain/user"

type UserApplicationService struct {
	userRepository UserRepositoryInterface
}

func NewUserApplicationService(userRepository UserRepositoryInterface) *UserApplicationService {
	return &UserApplicationService{
		userRepository: userRepository,
	}
}

func (s UserApplicationService) Regist(user domain.User) (domain.User, error) {
	savedUser, err := s.userRepository.Save(&user)
	return *savedUser, err
}