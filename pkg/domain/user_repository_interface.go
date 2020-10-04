package domain

type UserRepositoryInterface interface {
	Save(user *User) (*User, error)
	Find(email string) (*User, error)
	FindAll() ([]User, error)
	Begin()
	Rollback()
	Commit()
}
