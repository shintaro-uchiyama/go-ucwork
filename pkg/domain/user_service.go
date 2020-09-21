package domain

type UserService struct {
	userRepository UserRepositoryInterface
}

func NewUserService(userRepository UserRepositoryInterface) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u UserService) Exists(user User) bool{
	_, err := u.userRepository.Find(user.email)
	if err != nil {
		return false
	}
	return true
}

