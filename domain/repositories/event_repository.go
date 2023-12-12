package repositories

import "City-Pulse-API/domain/entities"

type EventRepository interface {
	AllEvents() ([]entities.Event, error)
	AllEventIDs() []string
	EventByID(id string) (*entities.Event, error)
	EventIDsForLocation(locationID string) ([]string, error)
	CreateEvent(city entities.Event) (entities.Event, error)
	UpdateEvent(id string, city entities.Event) (entities.Event, error)
	DeleteEvent(id string) (entities.Event, error)
}
