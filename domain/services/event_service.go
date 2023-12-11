package services

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/repositories"
)

type EventService struct {
	Repo         repositories.EventRepository
	LocationRepo repositories.LocationRepository
	CityRepo     repositories.CityRepository
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

func (service *EventService) EventByID(id string) (*EventDetails, error) {
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

func (service *EventService) DeleteEvent(id string) (entities.Event, error) {
	event, err := service.Repo.DeleteEvent(id)
	if err != nil {
		return entities.Event{}, err
	}
	return event, nil
}

func (service *EventService) UpdateEvent(id string, event entities.Event) (entities.Event, error) {
	event, err := service.Repo.UpdateEvent(id, event)
	if err != nil {
		return entities.Event{}, err
	}
	return event, nil
}
