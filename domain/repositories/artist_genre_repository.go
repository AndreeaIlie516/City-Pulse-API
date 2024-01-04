package repositories

import "City-Pulse-API/domain/entities"

type ArtistGenreRepository interface {
	AllArtistGenreAssociations() ([]entities.ArtistGenre, error)
	AllArtistGenreAssociationIDs() ([]uint, error)
	ArtistGenreAssociationByID(id uint) (*entities.ArtistGenre, error)
	ArtistGenreAssociation(artistID uint, genreID uint) (*entities.ArtistGenre, error)
	GenreIDsForArtist(artistID uint) ([]uint, error)
	ArtistIDsForGenre(genreID uint) ([]uint, error)
	CreateArtistGenreAssociation(artistGenreAssociation entities.ArtistGenre) (entities.ArtistGenre, error)
	UpdateArtistGenreAssociation(id uint, updatedArtistGenreAssociation entities.ArtistGenre) (entities.ArtistGenre, error)
	DeleteArtistGenreAssociation(id uint) (entities.ArtistGenre, error)
	DeleteGenreFromItsArtists(genreID uint) ([]entities.ArtistGenre, error)
	DeleteArtistFromItsGenres(artistID uint) ([]entities.ArtistGenre, error)
}
