package infrastructure

import (
	"errors"
	"github.com/google/uuid"
	domain "github.com/shintaro-uchiyama/go-ucwork/pkg/domain/user"
)

var _ domain.UserRepositoryInterface = (*InMemoryUserRepository)(nil)

type InMemoryUserRepository struct {
	db map[string]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		db: map[string]*domain.User{},
	}
}

func (r InMemoryUserRepository) Save(user *domain.User) (*domain.User, error) {
	user.SetUUID(uuid.New().String())
	r.db[user.GetUUID()] = user
	return user, nil
}

func (r InMemoryUserRepository) Find(email string) (*domain.User, error) {
	for _, record := range r.db {
		if record.GetEmail() == email {
			return record, nil
		}
	}
	return nil, errors.New("not found")
}
