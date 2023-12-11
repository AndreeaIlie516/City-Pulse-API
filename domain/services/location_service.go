package services

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/repositories"
)

type LocationService struct {
	Repo     repositories.LocationRepository
	CityRepo repositories.CityRepository
}

type LocationDetails struct {
	Location entities.Location
	City     entities.City
}

type LocationsByCity struct {
	City      entities.City
	Locations []entities.Location
}

func (service *LocationService) AllLocations() ([]entities.Location, error) {
	locations, err := service.Repo.AllLocations()
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (service *LocationService) LocationByID(id string) (*LocationDetails, error) {
	location, err := service.Repo.LocationByID(id)
	if err != nil {
		return nil, err
	}

	city, err := service.CityRepo.CityByID(location.CityID)
	if err != nil {
		return nil, err
	}

	locationDetails := &LocationDetails{
		Location: *location,
		City:     *city,
	}
	return locationDetails, nil
}

func (service *LocationService) LocationsByCityID(cityID string) (*LocationsByCity, error) {
	city, err := service.CityRepo.CityByID(cityID)
	if err != nil {
		return &LocationsByCity{}, err
	}

	locationIDs, err := service.Repo.LocationIDsForCity(cityID)
	var locations []entities.Location

	for _, locationID := range locationIDs {
		location, err := service.Repo.LocationByID(locationID)
		if err != nil {
			return &LocationsByCity{}, err
		}
		locations = append(locations, *location)
	}

	locationsByCity := &LocationsByCity{
		City:      *city,
		Locations: locations,
	}

	return locationsByCity, nil
}

func (service *LocationService) CreateLocation(location entities.Location) (entities.Location, error) {
	location, err := service.Repo.CreateLocation(location)
	if err != nil {
		return entities.Location{}, err
	}
	return location, nil
}

func (service *LocationService) DeleteLocation(id string) (entities.Location, error) {
	location, err := service.Repo.DeleteLocation(id)
	if err != nil {
		return entities.Location{}, err
	}
	return location, nil
}

func (service *LocationService) UpdateLocation(id string, location entities.Location) (entities.Location, error) {
	location, err := service.Repo.UpdateLocation(id, location)
	if err != nil {
		return entities.Location{}, err
	}
	return location, nil
}
