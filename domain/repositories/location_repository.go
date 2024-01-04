package repositories

import "City-Pulse-API/domain/entities"

type LocationRepository interface {
	AllLocations() ([]entities.Location, error)
	AllLocationIDs() ([]uint, error)
	LocationByID(id uint) (*entities.Location, error)
	LocationIDsForCity(cityID uint) ([]uint, error)
	CreateLocation(location entities.Location) (entities.Location, error)
	UpdateLocation(id uint, updatedLocation entities.Location) (entities.Location, error)
	DeleteLocation(id uint) (entities.Location, error)
}
