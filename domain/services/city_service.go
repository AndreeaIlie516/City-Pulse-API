package services

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/repositories"
)

type CityService struct {
	Repo repositories.CityRepository
}

func (service *CityService) AllCities() ([]entities.City, error) {
	cities, err := service.Repo.AllCities()
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func (service *CityService) CityByID(id string) (*entities.City, error) {
	city, err := service.Repo.CityByID(id)
	if err != nil {
		return nil, err
	}
	return city, nil
}

func (service *CityService) CreateCity(city entities.City) (entities.City, error) {
	city, err := service.Repo.CreateCity(city)
	if err != nil {
		return entities.City{}, err
	}
	return city, nil
}

func (service *CityService) DeleteCity(id string) (entities.City, error) {
	city, err := service.Repo.DeleteCity(id)
	if err != nil {
		return entities.City{}, err
	}
	return city, nil
}

func (service *CityService) UpdateCity(id string, city entities.City) (entities.City, error) {
	city, err := service.Repo.UpdateCity(id, city)
	if err != nil {
		return entities.City{}, err
	}
	return city, nil
}
