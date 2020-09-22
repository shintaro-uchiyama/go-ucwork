package domain

type UserService struct {
	userRepository UserRepositoryInterface
}

func NewUserService(userRepository UserRepositoryInterface) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u UserService) Exists(user User) bool {
	foundUser, err := u.userRepository.Find(user.GetEmail())
	if err != nil {
		return false
	}
	if foundUser != nil {
		return true
	}
	return false
}
