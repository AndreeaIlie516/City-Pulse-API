package repositories

import "City-Pulse-API/domain/entities"

type LocationRepository interface {
	AllLocations() ([]entities.Location, error)
	LocationByID(id string) (*entities.Location, error)
	LocationIDsForCity(cityID string) ([]string, error)
	CreateLocation(city entities.Location) (entities.Location, error)
	UpdateLocation(id string, location entities.Location) (entities.Location, error)
	DeleteLocation(id string) (entities.Location, error)
}
