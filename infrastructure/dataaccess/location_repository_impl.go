package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"errors"
	"strconv"
	"sync"
)

type InMemoryLocationRepository struct {
	locations []entities.Location
	mu        sync.RWMutex
}

func NewInMemoryLocationRepository() *InMemoryLocationRepository {
	return &InMemoryLocationRepository{}
}

func (r *InMemoryLocationRepository) AllLocations() ([]entities.Location, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.locations, nil
}

func (r *InMemoryLocationRepository) LocationByID(id string) (*entities.Location, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i, location := range r.locations {
		if location.ID == id {
			return &r.locations[i], nil
		}
	}

	return nil, errors.New("location not found")
}

func (r *InMemoryLocationRepository) LocationIDsForCity(cityID string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var locationIDs []string

	for _, location := range r.locations {
		if location.CityID == cityID {
			locationIDs = append(locationIDs, location.ID)
		}
	}

	return locationIDs, nil
}

func (r *InMemoryLocationRepository) CreateLocation(location entities.Location) (entities.Location, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	location.ID = strconv.Itoa(len(r.locations) + 1)
	r.locations = append(r.locations, location)
	return location, nil
}

func (r *InMemoryLocationRepository) DeleteLocation(id string) (entities.Location, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, location := range r.locations {
		if location.ID == id {
			r.locations = append(r.locations[:i], r.locations[i+1:]...)
			return location, nil
		}
	}
	return entities.Location{}, errors.New("location not found")
}

func (r *InMemoryLocationRepository) UpdateLocation(id string, updatedLocation entities.Location) (entities.Location, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, location := range r.locations {
		if location.ID == id {
			r.locations[i] = updatedLocation
			r.locations[i].ID = id
			return r.locations[i], nil
		}
	}

	return entities.Location{}, errors.New("location not found")
}
