package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"errors"
	"strconv"
	"sync"
)

type InMemoryEventRepository struct {
	events []entities.Event
	mu     sync.RWMutex
}

func NewInMemoryEventRepository() *InMemoryEventRepository {
	return &InMemoryEventRepository{}
}

func (r *InMemoryEventRepository) AllEvents() ([]entities.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.events, nil
}

func (r *InMemoryEventRepository) EventByID(id string) (*entities.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i, event := range r.events {
		if event.ID == id {
			return &r.events[i], nil
		}
	}

	return nil, errors.New("event not found")
}

func (r *InMemoryEventRepository) EventIDsForLocation(locationID string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var eventIDs []string

	for _, event := range r.events {
		if event.LocationID == locationID {
			eventIDs = append(eventIDs, event.ID)
		}
	}

	return eventIDs, nil
}

func (r *InMemoryEventRepository) EventIDsForCity(cityID string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *InMemoryEventRepository) CreateEvent(event entities.Event) (entities.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	event.ID = strconv.Itoa(len(r.events) + 1)
	r.events = append(r.events, event)
	return event, nil
}

func (r *InMemoryEventRepository) DeleteEvent(id string) (entities.Event, error) {
	for i, event := range r.events {
		if event.ID == id {
			r.events = append(r.events[:i], r.events[i+1:]...)
			return event, nil
		}
	}
	return entities.Event{}, errors.New("event not found")
}

func (r *InMemoryEventRepository) UpdateEvent(id string, updatedEvent entities.Event) (entities.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, event := range r.events {
		if event.ID == id {
			r.events[i] = updatedEvent
			r.events[i].ID = id
			return r.events[i], nil
		}
	}

	return entities.Event{}, errors.New("event not found")
}
