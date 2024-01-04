package services

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/repositories"
	"errors"
	"fmt"
)

type GenreService struct {
	Repo            repositories.GenreRepository
	ArtistGenreRepo repositories.ArtistGenreRepository
}

func (service *GenreService) AllGenres() ([]entities.Genre, error) {
	genres, err := service.Repo.AllGenres()
	if err != nil {
		return nil, err
	}
	return genres, nil
}

func (service *GenreService) GenreByID(idStr string) (*entities.Genre, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return nil, errors.New("invalid ID format")
	}

	genre, err := service.Repo.GenreByID(id)
	if err != nil {
		return nil, err
	}
	return genre, nil
}

func (service *GenreService) CreateGenre(genre entities.Genre) (entities.Genre, error) {
	genre, err := service.Repo.CreateGenre(genre)
	if err != nil {
		return entities.Genre{}, err
	}
	return genre, nil
}

func (service *GenreService) DeleteGenre(idStr string) (entities.Genre, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.Genre{}, errors.New("invalid ID format")
	}

	_, err := service.ArtistGenreRepo.DeleteGenreFromItsArtists(id)
	if err != nil {
		return entities.Genre{}, err
	}

	genre, err := service.Repo.DeleteGenre(id)
	if err != nil {
		return entities.Genre{}, err
	}
	return genre, nil
}

func (service *GenreService) UpdateGenre(idStr string, genre entities.Genre) (entities.Genre, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.Genre{}, errors.New("invalid ID format")
	}

	genre, err := service.Repo.UpdateGenre(id, genre)
	if err != nil {
		return entities.Genre{}, err
	}
	return genre, nil
}
