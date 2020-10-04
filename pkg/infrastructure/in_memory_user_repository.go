package infrastructure

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
)

var _ domain.UserRepositoryInterface = (*InMemoryUserRepository)(nil)

type InMemoryUserRepository struct {
	lock sync.RWMutex
	db   map[string]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		db: map[string]*domain.User{},
	}
}

func (r *InMemoryUserRepository) Save(user *domain.User) (*domain.User, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	user.SetUUID(uuid.New().String())
	r.db[user.UUID()] = user
	fmt.Println(fmt.Sprintf("db create user: %+v", user))
	return user, nil
}

func (r *InMemoryUserRepository) Find(email string) (*domain.User, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	for _, record := range r.db {
		if record.Email() == email {
			return record, nil
		}
	}
	return nil, nil
}

func (r *InMemoryUserRepository) FindAll() ([]domain.User, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	var users []domain.User
	for _, record := range r.db {
		fmt.Println(fmt.Sprintf("record: %+v", record))
		users = append(users, *record)
	}
	return users, nil
}
