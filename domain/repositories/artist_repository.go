package repositories

import "City-Pulse-API/domain/entities"

type ArtistRepository interface {
	AllArtists() ([]entities.Artist, error)
	AllArtistIDs() ([]uint, error)
	ArtistByID(id uint) (*entities.Artist, error)
	CreateArtist(artist entities.Artist) (entities.Artist, error)
	UpdateArtist(id uint, updatedArtist entities.Artist) (entities.Artist, error)
	DeleteArtist(id uint) (entities.Artist, error)
}
