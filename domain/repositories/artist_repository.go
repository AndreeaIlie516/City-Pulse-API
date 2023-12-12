package repositories

import "City-Pulse-API/domain/entities"

type ArtistRepository interface {
	AllArtists() ([]entities.Artist, error)
	AllArtistIDs() []string
	ArtistByID(id string) (*entities.Artist, error)
	CreateArtist(city entities.Artist) (entities.Artist, error)
	UpdateArtist(id string, artist entities.Artist) (entities.Artist, error)
	DeleteArtist(id string) (entities.Artist, error)
}
