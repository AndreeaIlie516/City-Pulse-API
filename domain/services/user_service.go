package services

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/repositories"
)

type UserService struct {
	Repo repositories.UserRepository
}

func (service *UserService) AllUsers() ([]entities.User, error) {
	users, err := service.Repo.AllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (service *UserService) UserByID(id string) (*entities.User, error) {
	user, err := service.Repo.UserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) CreateUser(user entities.User) (entities.User, error) {
	user, err := service.Repo.CreateUser(user)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (service *UserService) DeleteUser(id string) (entities.User, error) {
	user, err := service.Repo.DeleteUser(id)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (service *UserService) UpdateUser(id string, user entities.User) (entities.User, error) {
	user, err := service.Repo.UpdateUser(id, user)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}
