package repositories

import "City-Pulse-API/domain/entities"

type EventArtistRepository interface {
	AllEventArtistAssociations() ([]entities.EventArtist, error)
	AllEventArtistAssociationIDs() ([]uint, error)
	EventArtistAssociationByID(id uint) (*entities.EventArtist, error)
	EventArtistAssociation(artistID uint, genreID uint) (*entities.EventArtist, error)
	ArtistIDsForEvent(eventID uint) ([]uint, error)
	EventIDsForArtist(artistID uint) ([]uint, error)
	CreateEventArtistAssociation(eventArtistAssociation entities.EventArtist) (entities.EventArtist, error)
	UpdateEventArtistAssociation(id uint, updatedEventArtistAssociation entities.EventArtist) (entities.EventArtist, error)
	DeleteEventArtistAssociation(id uint) (entities.EventArtist, error)
	DeleteArtistFromItsEvents(artistID uint) ([]entities.EventArtist, error)
	DeleteEventFromItsArtists(eventID uint) ([]entities.EventArtist, error)
}
