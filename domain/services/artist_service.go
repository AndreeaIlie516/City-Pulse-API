package services

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/repositories"
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

func (service *ArtistService) ArtistByID(id string) (*entities.Artist, error) {
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

func (service *ArtistService) DeleteArtist(id string) (entities.Artist, error) {
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

func (service *ArtistService) UpdateArtist(id string, artist entities.Artist) (entities.Artist, error) {
	artist, err := service.Repo.UpdateArtist(id, artist)
	if err != nil {
		return entities.Artist{}, err
	}
	return artist, nil
}
