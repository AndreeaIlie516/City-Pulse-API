package repositories

import "City-Pulse-API/domain/entities"

type UserRepository interface {
	AllUsers() ([]entities.User, error)
	UserByID(id string) (*entities.User, error)
	CreateUser(user entities.User) (entities.User, error)
	UpdateUser(id string, user entities.User) (entities.User, error)
	DeleteUser(id string) (entities.User, error)
}
