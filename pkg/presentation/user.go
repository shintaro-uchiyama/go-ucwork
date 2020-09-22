package presentation

import "github.com/shintaro-uchiyama/go-ucwork/pkg/domain"

type User struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Users []User

type UserCreateRequest struct {
	User `json:"user" binding:"required"`
}

type UserListResponse struct {
	Users `json:"users" binding:"-"`
}

func NewUserListResponse(users domain.Users) *UserListResponse {
	var userListResponse UserListResponse
	for _, user := range users {
		userListResponse.Users = append(userListResponse.Users, User{
			Email: user.GetEmail(),
			Password: user.GetPassword(),
		})
	}
	return &userListResponse
}
