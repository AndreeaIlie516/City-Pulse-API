package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"errors"
	"strconv"
	"sync"
)

type InMemoryUserRepository struct {
	users []entities.User
	mu    sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{}
}

func (r *InMemoryUserRepository) AllUsers() ([]entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.users, nil
}

func (r *InMemoryUserRepository) UserByID(id string) (*entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i, user := range r.users {
		if user.ID == id {
			return &r.users[i], nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *InMemoryUserRepository) CreateUser(user entities.User) (entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user.ID = strconv.Itoa(len(r.users) + 1)
	r.users = append(r.users, user)
	return user, nil
}

func (r *InMemoryUserRepository) DeleteUser(id string) (entities.User, error) {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return user, nil
		}
	}
	return entities.User{}, errors.New("user not found")
}

func (r *InMemoryUserRepository) UpdateUser(id string, updatedUser entities.User) (entities.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, user := range r.users {
		if user.ID == id {
			r.users[i] = updatedUser
			r.users[i].ID = id
			return r.users[i], nil
		}
	}

	return entities.User{}, errors.New("user not found")
}
