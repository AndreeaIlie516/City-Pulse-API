package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/utils"
	"errors"
	"sync"
)

type InMemoryGenreRepository struct {
	genres []entities.Genre
	mu     sync.RWMutex
}

func NewInMemoryGenreRepository() *InMemoryGenreRepository {
	return &InMemoryGenreRepository{}
}

func (r *InMemoryGenreRepository) AllGenres() ([]entities.Genre, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.genres, nil
}

func (r *InMemoryGenreRepository) AllGenreIDs() []string {
	var genreIDs []string

	for _, genre := range r.genres {
		genreIDs = append(genreIDs, genre.ID)
	}

	return genreIDs
}

func (r *InMemoryGenreRepository) GenreByID(id string) (*entities.Genre, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i, genre := range r.genres {
		if genre.ID == id {
			return &r.genres[i], nil
		}
	}

	return nil, errors.New("genre not found")
}

func (r *InMemoryGenreRepository) CreateGenre(genre entities.Genre) (entities.Genre, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	genre.ID = utils.CreateUniqueID(utils.MinRange, utils.MaxRange, r.AllGenreIDs())
	r.genres = append(r.genres, genre)
	return genre, nil
}

func (r *InMemoryGenreRepository) DeleteGenre(id string) (entities.Genre, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, genre := range r.genres {
		if genre.ID == id {
			r.genres = append(r.genres[:i], r.genres[i+1:]...)
			return genre, nil
		}
	}
	return entities.Genre{}, errors.New("genre not found")
}

func (r *InMemoryGenreRepository) UpdateGenre(id string, updatedGenre entities.Genre) (entities.Genre, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, genre := range r.genres {
		if genre.ID == id {
			r.genres[i].Name = updatedGenre.Name
			r.genres[i].Description = updatedGenre.Description
			return r.genres[i], nil
		}
	}

	return entities.Genre{}, errors.New("genre not found")
}
