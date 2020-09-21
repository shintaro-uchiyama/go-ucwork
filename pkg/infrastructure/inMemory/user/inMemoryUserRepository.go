package infrastructure

import (
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