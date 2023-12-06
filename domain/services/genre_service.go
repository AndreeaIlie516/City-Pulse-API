package services

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/repositories"
)

type GenreService struct {
	Repo repositories.GenreRepository
}

func (service *GenreService) AllGenres() ([]entities.Genre, error) {
	genres, err := service.Repo.AllGenres()
	if err != nil {
		return nil, err
	}
	return genres, nil
}

func (service *GenreService) GenreByID(id string) (*entities.Genre, error) {
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

func (service *GenreService) DeleteGenre(id string) (entities.Genre, error) {
	genre, err := service.Repo.DeleteGenre(id)
	if err != nil {
		return entities.Genre{}, err
	}
	return genre, nil
}

func (service *GenreService) UpdateGenre(id string, genre entities.Genre) (entities.Genre, error) {
	genre, err := service.Repo.UpdateGenre(id, genre)
	if err != nil {
		return entities.Genre{}, err
	}
	return genre, nil
}
