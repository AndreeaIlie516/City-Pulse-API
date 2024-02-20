package repositories

import "City-Pulse-API/domain/entities"

type FavouriteEventRepository interface {
	AllFavouriteEventAssociations() ([]entities.FavouriteEvent, error)
	AllFavouriteEventAssociationsIDs() ([]uint, error)
	FavouriteEventAssociationByID(id uint) (*entities.FavouriteEvent, error)
	FavouriteEventAssociation(eventID uint, userID uint) (*entities.FavouriteEvent, error)
	UserIDsForEvent(eventID uint) ([]uint, error)
	EventIDsForUser(userID uint) ([]uint, error)
	AddEventToFavourites(favouriteEventAssociation entities.FavouriteEvent) (entities.FavouriteEvent, error)
	DeleteEventFromFavourites(id uint) (entities.FavouriteEvent, error)
	DeleteUserFromItsEvents(userID uint) ([]entities.FavouriteEvent, error)
	DeleteEventFromItsUsers(eventID uint) ([]entities.FavouriteEvent, error)
}
