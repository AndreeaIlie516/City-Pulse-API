package repositories

import "City-Pulse-API/domain/entities"

type CityRepository interface {
	AllCities() ([]entities.City, error)
	AllCityIDs() ([]uint, error)
	CityByID(id uint) (*entities.City, error)
	CreateCity(city entities.City) (entities.City, error)
	UpdateCity(id uint, updatedCity entities.City) (entities.City, error)
	DeleteCity(id uint) (entities.City, error)
}
