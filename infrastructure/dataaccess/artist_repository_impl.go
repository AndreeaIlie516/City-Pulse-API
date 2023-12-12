package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/utils"
	"errors"
	"sync"
)

type InMemoryArtistRepository struct {
	artists []entities.Artist
	mu      sync.RWMutex
}

func NewInMemoryArtistRepository() *InMemoryArtistRepository {
	return &InMemoryArtistRepository{}
}

func (r *InMemoryArtistRepository) AllArtists() ([]entities.Artist, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.artists, nil
}

func (r *InMemoryArtistRepository) AllArtistIDs() []string {
	var artistIDs []string

	for _, artist := range r.artists {
		artistIDs = append(artistIDs, artist.ID)
	}

	return artistIDs
}

func (r *InMemoryArtistRepository) ArtistByID(id string) (*entities.Artist, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i, artist := range r.artists {
		if artist.ID == id {
			return &r.artists[i], nil
		}
	}

	return nil, errors.New("artist not found")
}

func (r *InMemoryArtistRepository) CreateArtist(artist entities.Artist) (entities.Artist, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	artistIDs := r.AllArtistIDs()
	artist.ID = utils.CreateUniqueID(utils.MinRange, utils.MaxRange, artistIDs)
	r.artists = append(r.artists, artist)
	return artist, nil
}

func (r *InMemoryArtistRepository) DeleteArtist(id string) (entities.Artist, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, artist := range r.artists {
		if artist.ID == id {
			r.artists = append(r.artists[:i], r.artists[i+1:]...)
			return artist, nil
		}
	}
	return entities.Artist{}, errors.New("artist not found")
}

func (r *InMemoryArtistRepository) UpdateArtist(id string, updatedArtist entities.Artist) (entities.Artist, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, artist := range r.artists {
		if artist.ID == id {
			r.artists[i] = updatedArtist
			r.artists[i].ID = id
			return r.artists[i], nil
		}
	}

	return entities.Artist{}, errors.New("artist not found")
}
