package services

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/repositories"
	"errors"
	"fmt"
)

type EventService struct {
	Repo            repositories.EventRepository
	LocationRepo    repositories.LocationRepository
	CityRepo        repositories.CityRepository
	EventArtistRepo repositories.EventArtistRepository
}

type EventDetails struct {
	Event    entities.Event
	Location entities.Location
	City     entities.City
}

func (service *EventService) AllEvents() ([]entities.Event, error) {
	events, err := service.Repo.AllEvents()
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (service *EventService) EventByID(idStr string) (*EventDetails, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return nil, errors.New("invalid ID format")
	}

	event, err := service.Repo.EventByID(id)
	if err != nil {
		return nil, err
	}

	location, err := service.LocationRepo.LocationByID(event.LocationID)
	if err != nil {
		return nil, err
	}

	city, err := service.CityRepo.CityByID(location.CityID)
	if err != nil {
		return nil, err
	}

	eventDetails := &EventDetails{
		Event:    *event,
		Location: *location,
		City:     *city,
	}
	return eventDetails, nil
}

func (service *EventService) CreateEvent(event entities.Event) (entities.Event, error) {
	_, err := service.LocationRepo.LocationByID(event.LocationID)
	if err != nil {
		return entities.Event{}, err
	}

	event, err = service.Repo.CreateEvent(event)
	if err != nil {
		return entities.Event{}, err
	}
	return event, nil
}

func (service *EventService) DeleteEvent(idStr string) (entities.Event, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.Event{}, errors.New("invalid ID format")
	}

	_, err := service.EventArtistRepo.DeleteEventFromItsArtists(id)
	if err != nil {
		return entities.Event{}, err
	}
	event, err := service.Repo.DeleteEvent(id)
	if err != nil {
		return entities.Event{}, err
	}
	return event, nil
}

func (service *EventService) UpdateEvent(idStr string, event entities.Event) (entities.Event, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.Event{}, errors.New("invalid ID format")
	}

	event, err := service.Repo.UpdateEvent(id, event)
	if err != nil {
		return entities.Event{}, err
	}
	return event, nil
}
