package user

import (
	"errors"
	"sync"
)

type InMemoryUserRepository struct {
	data map[string]*User
	mu   sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		data: make(map[string]*User),
	}
}

func (r *InMemoryUserRepository) GetUserByID(id string) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.data[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *InMemoryUserRepository) CreateUser(user *User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[user.ID]; exists {
		return errors.New("user already exists")
	}

	r.data[user.ID] = user
	return nil
}