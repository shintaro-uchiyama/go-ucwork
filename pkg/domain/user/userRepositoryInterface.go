package domain

type UserRepositoryInterface interface {
	Save(user *User) (*User, error)
}