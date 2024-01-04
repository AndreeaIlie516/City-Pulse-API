package services

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/repositories"
	"errors"
	"fmt"
)

type CityService struct {
	Repo         repositories.CityRepository
	LocationRepo repositories.LocationRepository
}

func (service *CityService) AllCities() ([]entities.City, error) {
	cities, err := service.Repo.AllCities()
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func (service *CityService) CityByID(idStr string) (*entities.City, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return nil, errors.New("invalid ID format")
	}

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

func (service *CityService) DeleteCity(idStr string) (entities.City, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.City{}, errors.New("invalid ID format")
	}

	locations, err := service.LocationRepo.LocationIDsForCity(id)
	if err != nil {
		return entities.City{}, err
	}

	if len(locations) > 0 {
		return entities.City{}, errors.New("cannot delete city with associated locations")
	}
	city, err := service.Repo.DeleteCity(id)
	if err != nil {
		return entities.City{}, err
	}
	return city, nil
}

func (service *CityService) UpdateCity(idStr string, city entities.City) (entities.City, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.City{}, errors.New("invalid ID format")
	}

	city, err := service.Repo.UpdateCity(id, city)
	if err != nil {
		return entities.City{}, err
	}
	return city, nil
}
