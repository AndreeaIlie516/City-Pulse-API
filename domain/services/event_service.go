package services

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/repositories"
	"errors"
	"fmt"
)

type EventService struct {
	Repo               repositories.EventRepository
	LocationRepo       repositories.LocationRepository
	CityRepo           repositories.CityRepository
	EventArtistRepo    repositories.EventArtistRepository
	FavouriteEventRepo repositories.FavouriteEventRepository
}

type EventDetails struct {
	Event    entities.Event
	Location entities.Location
	City     entities.City
}

type EventWithLocation struct {
	Event    entities.Event
	Location entities.Location
}

type EventsByLocation struct {
	Location entities.Location
	Events   []entities.Event
}

type EventsByCity struct {
	City   entities.City
	Events []EventWithLocation
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

func (service *EventService) EventsByLocationID(locationIDStr string) (*EventsByLocation, error) {
	var locationID uint
	if _, err := fmt.Sscanf(locationIDStr, "%d", &locationID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	location, err := service.LocationRepo.LocationByID(locationID)
	if err != nil {
		return &EventsByLocation{}, err
	}

	eventIDs, err := service.Repo.EventIDsForLocation(locationID)
	var events []entities.Event

	for _, eventID := range eventIDs {
		event, err := service.Repo.EventByID(eventID)
		if err != nil {
			return &EventsByLocation{}, err
		}
		events = append(events, *event)
	}

	eventsByLocation := &EventsByLocation{
		Location: *location,
		Events:   events,
	}

	return eventsByLocation, nil
}

func (service *EventService) EventsByCityID(cityIDStr string) (*EventsByCity, error) {
	var cityID uint
	if _, err := fmt.Sscanf(cityIDStr, "%d", &cityID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	city, err := service.CityRepo.CityByID(cityID)
	if err != nil {
		return &EventsByCity{}, err
	}

	eventIDs, err := service.Repo.EventIDsForCity(cityID)
	var events []EventWithLocation

	for _, eventID := range eventIDs {
		event, err := service.Repo.EventByID(eventID)
		if err != nil {
			return &EventsByCity{}, err
		}
		location, _ := service.LocationRepo.LocationByID(event.LocationID)
		eventWithLocation := &EventWithLocation{
			Event:    *event,
			Location: *location,
		}
		events = append(events, *eventWithLocation)
	}

	eventsByCity := &EventsByCity{
		City:   *city,
		Events: events,
	}

	return eventsByCity, nil
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

	_, err = service.FavouriteEventRepo.DeleteEventFromItsUsers(id)
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
