package presentation

import (
	"fmt"
	domain "github.com/shintaro-uchiyama/go-ucwork/pkg/domain/user"
)

type UserController struct {
	userApplicationService UserApplicationServiceInterface
}

func NewUserController(service UserApplicationServiceInterface) *UserController {
	return &UserController{
		userApplicationService: service,
	}
}

func (c UserController) Regist() (domain.User, error) {
	fmt.Print("Please Enter email: ")
	var email string
	_, emailErr := fmt.Scanln(&email)
	if emailErr != nil {
		return domain.User{}, emailErr
	}

	fmt.Print("Please Enter password: ")
	var password string
	_, passwordErr := fmt.Scanln(&password)
	if passwordErr != nil {
		return domain.User{}, emailErr
	}

	requestUser := domain.NewUser(email, password)
	user, err := c.userApplicationService.Regist(*requestUser)
	if err != nil {
		return domain.User{}, err
	}
	return user, err
}
