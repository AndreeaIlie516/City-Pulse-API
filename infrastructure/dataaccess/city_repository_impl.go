package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/utils"
	"errors"
	"sync"
)

type InMemoryCityRepository struct {
	cities []entities.City
	mu     sync.RWMutex
}

func NewInMemoryCityRepository() *InMemoryCityRepository {
	return &InMemoryCityRepository{}
}

func (r *InMemoryCityRepository) AllCities() ([]entities.City, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.cities, nil
}

func (r *InMemoryCityRepository) AllCityIDs() []string {
	var cityIDs []string

	for _, city := range r.cities {
		cityIDs = append(cityIDs, city.ID)
	}

	return cityIDs
}

func (r *InMemoryCityRepository) CityByID(id string) (*entities.City, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i, city := range r.cities {
		if city.ID == id {
			return &r.cities[i], nil
		}
	}

	return nil, errors.New("city not found")
}

func (r *InMemoryCityRepository) CreateCity(city entities.City) (entities.City, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	city.ID = utils.CreateUniqueID(utils.MinRange, utils.MaxRange, r.AllCityIDs())
	r.cities = append(r.cities, city)
	return city, nil
}

func (r *InMemoryCityRepository) DeleteCity(id string) (entities.City, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, city := range r.cities {
		if city.ID == id {
			r.cities = append(r.cities[:i], r.cities[i+1:]...)
			return city, nil
		}
	}
	return entities.City{}, errors.New("city not found")
}

func (r *InMemoryCityRepository) UpdateCity(id string, updatedCity entities.City) (entities.City, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, city := range r.cities {
		if city.ID == id {
			r.cities[i] = updatedCity
			r.cities[i].ID = id
			return r.cities[i], nil
		}
	}

	return entities.City{}, errors.New("city not found")
}
