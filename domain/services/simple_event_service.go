package services

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/repositories"
	"errors"
	"fmt"
)

type SimpleEventService struct {
	Repo repositories.SimpleEventRepository
}

func (service *SimpleEventService) AllEvents() ([]entities.SimpleEvent, error) {
	events, err := service.Repo.AllEvents()
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (service *SimpleEventService) EventByID(idStr string) (*entities.SimpleEvent, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return nil, errors.New("invalid ID format")
	}

	event, err := service.Repo.EventByID(id)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (service *SimpleEventService) CreateEvent(event entities.SimpleEvent) (entities.SimpleEvent, error) {
	event, err := service.Repo.CreateEvent(event)
	if err != nil {
		return entities.SimpleEvent{}, err
	}
	return event, nil
}

func (service *SimpleEventService) DeleteEvent(idStr string) (entities.SimpleEvent, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.SimpleEvent{}, errors.New("invalid ID format")
	}

	event, err := service.Repo.DeleteEvent(id)
	if err != nil {
		return entities.SimpleEvent{}, err
	}
	return event, nil
}

func (service *SimpleEventService) UpdateEvent(idStr string, event entities.SimpleEvent) (entities.SimpleEvent, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.SimpleEvent{}, errors.New("invalid ID format")
	}

	event, err := service.Repo.UpdateEvent(id, event)
	if err != nil {
		return entities.SimpleEvent{}, err
	}
	return event, nil
}

func (service *SimpleEventService) FavouriteEvents() ([]entities.SimpleEvent, error) {
	favouriteEvents, err := service.Repo.FavouriteEvents()
	if err != nil {
		return nil, err
	}
	return favouriteEvents, nil
}

func (service *SimpleEventService) PrivateEvents() ([]entities.SimpleEvent, error) {
	privateEvents, err := service.Repo.PrivateEvents()
	if err != nil {
		return nil, err
	}
	return privateEvents, nil
}

func (service *SimpleEventService) AddEventToFavourites(idStr string) (entities.SimpleEvent, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.SimpleEvent{}, errors.New("invalid ID format")
	}

	event, err := service.Repo.AddEventToFavourites(id)
	if err != nil {
		return entities.SimpleEvent{}, err
	}
	return event, nil
}

func (service *SimpleEventService) DeleteEventFromFavourites(idStr string) (entities.SimpleEvent, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.SimpleEvent{}, errors.New("invalid ID format")
	}

	event, err := service.Repo.DeleteEventFromFavourites(id)
	if err != nil {
		return entities.SimpleEvent{}, err
	}
	return event, nil
}
