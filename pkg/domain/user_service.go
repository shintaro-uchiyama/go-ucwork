package domain

type UserService struct {
	userRepository UserRepositoryInterface
}

func NewUserService(userRepository UserRepositoryInterface) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u UserService) Exists(user User) (bool, error) {
	foundUser, err := u.userRepository.Find(user.Email())
	if err != nil {
		return false, err
	}
	if foundUser != nil {
		return true, nil
	}
	return false, nil
}
