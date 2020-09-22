package infrastructure

import (
	"github.com/google/uuid"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
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

func (r *InMemoryUserRepository) Save(user *domain.User) error {
	user.SetUUID(uuid.New().String())
	r.db[user.UUID()] = user
	return nil
}

func (r *InMemoryUserRepository) Find(email string) (*domain.User, error) {
	for _, record := range r.db {
		if record.Email() == email {
			return record, nil
		}
	}
	return nil, nil
}

func (r *InMemoryUserRepository) FindAll() (domain.Users, error) {
	var users domain.Users
	for _, record := range r.db {
		users = append(users, *record)
	}
	return users, nil
}
