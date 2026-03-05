package repository

import (
	"errors"
	"sync"
	"github.com/Ipnu22Aris33/user-service-go/internal/domain"
)

type InMemoryUserRepository struct {
	data map[string]*domain.User
	mu   sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		data: make(map[string]*domain.User),
	}
}

func (r *InMemoryUserRepository) GetUserByID(id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.data[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *InMemoryUserRepository) CreateUser(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[user.ID]; exists {
		return errors.New("user already exists")
	}

	r.data[user.ID] = user
	return nil
}