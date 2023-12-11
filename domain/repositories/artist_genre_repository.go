package repositories

import "City-Pulse-API/domain/entities"

type ArtistGenreRepository interface {
	AllArtistGenreAssociations() ([]entities.ArtistGenre, error)
	ArtistGenreAssociationByID(id string) (*entities.ArtistGenre, error)
	ArtistGenreAssociation(artistID string, genreID string) (*entities.ArtistGenre, error)
	GenreIDsForArtist(artistID string) ([]string, error)
	ArtistIDsForGenre(genreID string) ([]string, error)
	CreateArtistGenreAssociation(artistGenreAssociation entities.ArtistGenre) (entities.ArtistGenre, error)
	UpdateArtistGenreAssociation(id string, updatedArtistGenreAssociation entities.ArtistGenre) (entities.ArtistGenre, error)
	DeleteArtistGenreAssociation(id string) (entities.ArtistGenre, error)
	DeleteGenreFromItsArtists(genreID string) ([]entities.ArtistGenre, error)
	DeleteArtistFromItsGenres(artistID string) ([]entities.ArtistGenre, error)
}
