package services

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/repositories"
	"errors"
	"fmt"
)

type ArtistService struct {
	Repo            repositories.ArtistRepository
	ArtistGenreRepo repositories.ArtistGenreRepository
	EventArtistRepo repositories.EventArtistRepository
}

func (service *ArtistService) AllArtists() ([]entities.Artist, error) {
	artists, err := service.Repo.AllArtists()
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func (service *ArtistService) ArtistByID(idStr string) (*entities.Artist, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return nil, errors.New("invalid ID format")
	}

	artist, err := service.Repo.ArtistByID(id)
	if err != nil {
		return nil, err
	}
	return artist, nil
}

func (service *ArtistService) CreateArtist(artist entities.Artist) (entities.Artist, error) {
	artist, err := service.Repo.CreateArtist(artist)
	if err != nil {
		return entities.Artist{}, err
	}
	return artist, nil
}

func (service *ArtistService) DeleteArtist(idStr string) (entities.Artist, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.Artist{}, errors.New("invalid ID format")
	}

	_, err := service.ArtistGenreRepo.DeleteArtistFromItsGenres(id)
	if err != nil {
		return entities.Artist{}, err
	}

	_, err = service.EventArtistRepo.DeleteArtistFromItsEvents(id)
	if err != nil {
		return entities.Artist{}, err
	}

	artist, err := service.Repo.DeleteArtist(id)
	if err != nil {
		return entities.Artist{}, err
	}
	return artist, nil
}

func (service *ArtistService) UpdateArtist(idStr string, artist entities.Artist) (entities.Artist, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.Artist{}, errors.New("invalid ID format")
	}

	artist, err := service.Repo.UpdateArtist(id, artist)
	if err != nil {
		return entities.Artist{}, err
	}
	return artist, nil
}
