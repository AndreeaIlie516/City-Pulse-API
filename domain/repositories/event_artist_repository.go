package repositories

import "City-Pulse-API/domain/entities"

type EventArtistRepository interface {
	AllEventArtistAssociations() ([]entities.EventArtist, error)
	AllEventArtistAssociationIDs() []string
	EventArtistAssociationByID(id string) (*entities.EventArtist, error)
	EventArtistAssociation(artistID string, genreID string) (*entities.EventArtist, error)
	ArtistIDsForEvent(eventID string) ([]string, error)
	EventIDsForArtist(artistID string) ([]string, error)
	CreateEventArtistAssociation(eventArtistAssociation entities.EventArtist) (entities.EventArtist, error)
	UpdateEventArtistAssociation(id string, updatedEventArtistAssociation entities.EventArtist) (entities.EventArtist, error)
	DeleteEventArtistAssociation(id string) (entities.EventArtist, error)
	DeleteArtistFromItsEvents(artistID string) ([]entities.EventArtist, error)
	DeleteEventFromItsArtists(eventID string) ([]entities.EventArtist, error)
}
