package services

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/repositories"
	"errors"
	"fmt"
)

type ArtistGenreService struct {
	Repo       repositories.ArtistGenreRepository
	ArtistRepo repositories.ArtistRepository
	GenreRepo  repositories.GenreRepository
}

type ArtistGenreDetail struct {
	Association entities.ArtistGenre
	Artist      entities.Artist
	Genre       entities.Genre
}

type ArtistWithGenres struct {
	Artist entities.Artist
	Genres []entities.Genre
}

type GenreWithArtists struct {
	Genre   entities.Genre
	Artists []entities.Artist
}

func (service *ArtistGenreService) AllArtistGenreAssociations() ([]entities.ArtistGenre, error) {
	artistGenreAssociations, err := service.Repo.AllArtistGenreAssociations()
	if err != nil {
		return nil, err
	}
	return artistGenreAssociations, nil
}

func (service *ArtistGenreService) ArtistGenreAssociationByID(idStr string) (*ArtistGenreDetail, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return nil, errors.New("invalid ID format")
	}

	artistGenreAssociation, err := service.Repo.ArtistGenreAssociationByID(id)
	if err != nil {
		return nil, err
	}

	artist, err := service.ArtistRepo.ArtistByID(artistGenreAssociation.ArtistID)
	if err != nil {
		return nil, err
	}

	genre, err := service.GenreRepo.GenreByID(artistGenreAssociation.GenreID)
	if err != nil {
		return nil, err
	}

	artistGenreDetail := &ArtistGenreDetail{
		Association: *artistGenreAssociation,
		Artist:      *artist,
		Genre:       *genre,
	}
	return artistGenreDetail, nil
}

func (service *ArtistGenreService) ArtistGenreAssociation(artistIDStr string, genreIDStr string) (*ArtistGenreDetail, error) {
	var artistID uint
	if _, err := fmt.Sscanf(artistIDStr, "%d", &artistID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	artist, err := service.ArtistRepo.ArtistByID(artistID)
	if err != nil {
		return nil, err
	}

	var genreID uint
	if _, err := fmt.Sscanf(genreIDStr, "%d", &genreID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	genre, err := service.GenreRepo.GenreByID(genreID)
	if err != nil {
		return nil, err
	}

	artistGenreAssociation, err := service.Repo.ArtistGenreAssociation(artistID, genreID)
	if err != nil {
		return nil, err
	}

	artistGenreDetail := &ArtistGenreDetail{
		Association: *artistGenreAssociation,
		Artist:      *artist,
		Genre:       *genre,
	}
	return artistGenreDetail, nil
}

func (service *ArtistGenreService) ArtistWithGenres(artistIDStr string) (*ArtistWithGenres, error) {
	var artistID uint
	if _, err := fmt.Sscanf(artistIDStr, "%d", &artistID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	artist, err := service.ArtistRepo.ArtistByID(artistID)
	if err != nil {
		return &ArtistWithGenres{}, err
	}

	genreIDs, err := service.Repo.GenreIDsForArtist(artistID)
	var genres []entities.Genre

	for _, genreID := range genreIDs {
		genre, err := service.GenreRepo.GenreByID(genreID)
		if err != nil {
			return &ArtistWithGenres{}, err
		}
		genres = append(genres, *genre)
	}

	artistWithGenres := &ArtistWithGenres{
		Artist: *artist,
		Genres: genres,
	}

	return artistWithGenres, nil
}

func (service *ArtistGenreService) GenreWithArtists(genreIDStr string) (*GenreWithArtists, error) {
	var genreID uint
	if _, err := fmt.Sscanf(genreIDStr, "%d", &genreID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	genre, err := service.GenreRepo.GenreByID(genreID)
	if err != nil {
		return &GenreWithArtists{}, err
	}

	artistIDs, err := service.Repo.ArtistIDsForGenre(genreID)
	var artists []entities.Artist

	for _, artistID := range artistIDs {
		artist, err := service.ArtistRepo.ArtistByID(artistID)
		if err != nil {
			return &GenreWithArtists{}, err
		}
		artists = append(artists, *artist)
	}

	genreWithArtists := &GenreWithArtists{
		Genre:   *genre,
		Artists: artists,
	}

	return genreWithArtists, nil
}

func (service *ArtistGenreService) CreateArtistGenreAssociation(artistGenreAssociation entities.ArtistGenre) (entities.ArtistGenre, error) {
	_, err := service.GenreRepo.GenreByID(artistGenreAssociation.GenreID)
	if err != nil {
		return entities.ArtistGenre{}, err
	}

	_, err = service.ArtistRepo.ArtistByID(artistGenreAssociation.ArtistID)
	if err != nil {
		return entities.ArtistGenre{}, err
	}

	artistGenreAssociation, err = service.Repo.CreateArtistGenreAssociation(artistGenreAssociation)
	if err != nil {
		return entities.ArtistGenre{}, err
	}
	return artistGenreAssociation, nil
}

func (service *ArtistGenreService) DeleteArtistGenreAssociation(idStr string) (entities.ArtistGenre, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.ArtistGenre{}, errors.New("invalid ID format")
	}

	artistGenreAssociation, err := service.Repo.DeleteArtistGenreAssociation(id)
	if err != nil {
		return entities.ArtistGenre{}, err
	}
	return artistGenreAssociation, nil
}

func (service *ArtistGenreService) DeleteGenreFromItsArtists(genreIDStr string) ([]entities.ArtistGenre, error) {
	var genreID uint
	if _, err := fmt.Sscanf(genreIDStr, "%d", &genreID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	_, err := service.GenreRepo.GenreByID(genreID)
	if err != nil {
		return []entities.ArtistGenre{}, err
	}

	artistGenreAssociation, err := service.Repo.DeleteGenreFromItsArtists(genreID)
	if err != nil {
		return []entities.ArtistGenre{}, err
	}
	return artistGenreAssociation, nil
}

func (service *ArtistGenreService) DeleteArtistFromItsGenres(artistIDStr string) ([]entities.ArtistGenre, error) {
	var artistID uint
	if _, err := fmt.Sscanf(artistIDStr, "%d", &artistID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	_, err := service.ArtistRepo.ArtistByID(artistID)
	if err != nil {
		return []entities.ArtistGenre{}, err
	}

	artistGenreAssociation, err := service.Repo.DeleteArtistFromItsGenres(artistID)
	if err != nil {
		return []entities.ArtistGenre{}, err
	}
	return artistGenreAssociation, nil
}

func (service *ArtistGenreService) UpdateArtistGenreAssociation(idStr string, artistGenreAssociation entities.ArtistGenre) (entities.ArtistGenre, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.ArtistGenre{}, errors.New("invalid ID format")
	}

	artistGenreAssociation, err := service.Repo.UpdateArtistGenreAssociation(id, artistGenreAssociation)
	if err != nil {
		return entities.ArtistGenre{}, err
	}
	return artistGenreAssociation, nil
}
