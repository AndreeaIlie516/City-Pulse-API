package repositories

import "City-Pulse-API/domain/entities"

type EventRepository interface {
	AllEvents() ([]entities.Event, error)
	AllEventIDs() ([]uint, error)
	EventByID(id uint) (*entities.Event, error)
	EventIDsForLocation(locationID uint) ([]uint, error)
	CreateEvent(city entities.Event) (entities.Event, error)
	UpdateEvent(id uint, updatedEvent entities.Event) (entities.Event, error)
	DeleteEvent(id uint) (entities.Event, error)
}
