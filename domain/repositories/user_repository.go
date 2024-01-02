package repositories

import "City-Pulse-API/domain/entities"

type UserRepository interface {
	AllUsers() ([]entities.User, error)
	UserByID(id uint) (*entities.User, error)
	AllUserIDs() ([]uint, error)
	CreateUser(user entities.User) (entities.User, error)
	UpdateUser(id uint, user entities.User) (entities.User, error)
	DeleteUser(id uint) (entities.User, error)
}
