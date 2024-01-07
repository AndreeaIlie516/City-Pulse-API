package repositories

import "City-Pulse-API/domain/entities"

type SimpleEventRepository interface {
	AllEvents() ([]entities.SimpleEvent, error)
	AllEventIDs() ([]uint, error)
	EventByID(id uint) (*entities.SimpleEvent, error)
	CreateEvent(city entities.SimpleEvent) (entities.SimpleEvent, error)
	UpdateEvent(id uint, updatedEvent entities.SimpleEvent) (entities.SimpleEvent, error)
	DeleteEvent(id uint) (entities.SimpleEvent, error)
	FavouriteEvents() ([]entities.SimpleEvent, error)
	PrivateEvents() ([]entities.SimpleEvent, error)
	AddEventToFavourites(id uint) (entities.SimpleEvent, error)
	DeleteEventFromFavourites(id uint) (entities.SimpleEvent, error)
}
