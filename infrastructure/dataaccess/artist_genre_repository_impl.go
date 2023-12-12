package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/utils"
	"errors"
	"sync"
)

type InMemoryArtistGenreRepository struct {
	artistGenreAssociations []entities.ArtistGenre
	mu                      sync.RWMutex
}

func NewInMemoryArtistGenreRepository() *InMemoryArtistGenreRepository {
	return &InMemoryArtistGenreRepository{}
}

func (r *InMemoryArtistGenreRepository) AllArtistGenreAssociations() ([]entities.ArtistGenre, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.artistGenreAssociations, nil
}

func (r *InMemoryArtistGenreRepository) AllArtistGenreAssociationIDs() []string {
	var artistGenreAssociationIDs []string

	for _, artistGenreAssociation := range r.artistGenreAssociations {
		artistGenreAssociationIDs = append(artistGenreAssociationIDs, artistGenreAssociation.ID)
	}

	return artistGenreAssociationIDs
}

func (r *InMemoryArtistGenreRepository) ArtistGenreAssociationByID(id string) (*entities.ArtistGenre, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i, artistGenreAssociation := range r.artistGenreAssociations {
		if artistGenreAssociation.ID == id {
			return &r.artistGenreAssociations[i], nil
		}
	}

	return nil, errors.New("artist genre association not found")
}

func (r *InMemoryArtistGenreRepository) ArtistGenreAssociation(artistID string, genreID string) (*entities.ArtistGenre, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i, artistGenreAssociation := range r.artistGenreAssociations {
		if artistGenreAssociation.ArtistID == artistID && artistGenreAssociation.GenreID == genreID {
			return &r.artistGenreAssociations[i], nil
		}
	}

	return nil, errors.New("artist genre association not found")
}

func (r *InMemoryArtistGenreRepository) GenreIDsForArtist(artistID string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var genreIDs []string

	for _, artistGenreAssociation := range r.artistGenreAssociations {
		if artistGenreAssociation.ArtistID == artistID {
			genreIDs = append(genreIDs, artistGenreAssociation.GenreID)
		}
	}

	return genreIDs, nil
}

func (r *InMemoryArtistGenreRepository) ArtistIDsForGenre(genreID string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var artistIDs []string

	for _, artistGenreAssociation := range r.artistGenreAssociations {
		if artistGenreAssociation.GenreID == genreID {
			artistIDs = append(artistIDs, artistGenreAssociation.ArtistID)
		}
	}

	return artistIDs, nil
}

func (r *InMemoryArtistGenreRepository) CreateArtistGenreAssociation(artistGenreAssociation entities.ArtistGenre) (entities.ArtistGenre, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	artistGenreAssociation.ID = utils.CreateUniqueID(utils.MinRange, utils.MaxRange, r.AllArtistGenreAssociationIDs())
	r.artistGenreAssociations = append(r.artistGenreAssociations, artistGenreAssociation)
	return artistGenreAssociation, nil
}

func (r *InMemoryArtistGenreRepository) DeleteArtistGenreAssociation(id string) (entities.ArtistGenre, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, artistGenreAssociation := range r.artistGenreAssociations {
		if artistGenreAssociation.ID == id {
			r.artistGenreAssociations = append(r.artistGenreAssociations[:i], r.artistGenreAssociations[i+1:]...)
			return artistGenreAssociation, nil
		}
	}
	return entities.ArtistGenre{}, errors.New("artist genre associations not found")
}

func (r *InMemoryArtistGenreRepository) DeleteGenreFromItsArtists(genreID string) ([]entities.ArtistGenre, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var remainingAssociations []entities.ArtistGenre
	var deletedAssociations []entities.ArtistGenre
	for _, association := range r.artistGenreAssociations {
		if association.GenreID == genreID {
			deletedAssociations = append(deletedAssociations, association)
		} else {
			remainingAssociations = append(remainingAssociations, association)
		}
	}
	r.artistGenreAssociations = remainingAssociations
	return deletedAssociations, nil
}

func (r *InMemoryArtistGenreRepository) DeleteArtistFromItsGenres(artistID string) ([]entities.ArtistGenre, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var remainingAssociations []entities.ArtistGenre
	var deletedAssociations []entities.ArtistGenre
	for _, association := range r.artistGenreAssociations {
		if association.ArtistID == artistID {
			deletedAssociations = append(deletedAssociations, association)
		} else {
			remainingAssociations = append(remainingAssociations, association)
		}
	}
	r.artistGenreAssociations = remainingAssociations
	return deletedAssociations, nil
}

func (r *InMemoryArtistGenreRepository) UpdateArtistGenreAssociation(id string, updatedArtistGenreAssociation entities.ArtistGenre) (entities.ArtistGenre, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, artistGenreAssociation := range r.artistGenreAssociations {
		if artistGenreAssociation.ID == id {
			r.artistGenreAssociations[i] = updatedArtistGenreAssociation
			r.artistGenreAssociations[i].ID = id
			return r.artistGenreAssociations[i], nil
		}
	}

	return entities.ArtistGenre{}, errors.New("artist genre association not found")
}
