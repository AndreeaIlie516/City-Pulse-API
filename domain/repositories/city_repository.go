package repositories

import "City-Pulse-API/domain/entities"

type CityRepository interface {
	AllCities() ([]entities.City, error)
	AllCityIDs() []string
	CityByID(id string) (*entities.City, error)
	CreateCity(city entities.City) (entities.City, error)
	UpdateCity(id string, city entities.City) (entities.City, error)
	DeleteCity(id string) (entities.City, error)
}
