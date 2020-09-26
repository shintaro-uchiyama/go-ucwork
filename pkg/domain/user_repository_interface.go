package domain

type UserRepositoryInterface interface {
	Save(user *User) error
	Find(email string) (*User, error)
	FindAll() (Users, error)
}
