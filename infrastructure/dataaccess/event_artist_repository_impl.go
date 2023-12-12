package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"errors"
	"strconv"
	"sync"
)

type InMemoryEventArtistRepository struct {
	eventArtistAssociations []entities.EventArtist
	mu                      sync.RWMutex
}

func NewInMemoryEventArtistRepository() *InMemoryEventArtistRepository {
	return &InMemoryEventArtistRepository{}
}

func (r *InMemoryEventArtistRepository) AllEventArtistAssociations() ([]entities.EventArtist, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.eventArtistAssociations, nil
}

func (r *InMemoryEventArtistRepository) EventArtistAssociationByID(id string) (*entities.EventArtist, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i, eventArtistAssociation := range r.eventArtistAssociations {
		if eventArtistAssociation.ID == id {
			return &r.eventArtistAssociations[i], nil
		}
	}

	return nil, errors.New("event artist association not found")
}

func (r *InMemoryEventArtistRepository) EventArtistAssociation(eventID string, artistID string) (*entities.EventArtist, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i, eventArtistAssociation := range r.eventArtistAssociations {
		if eventArtistAssociation.EventID == eventID && eventArtistAssociation.ArtistID == artistID {
			return &r.eventArtistAssociations[i], nil
		}
	}

	return nil, errors.New("event artist association not found")
}

func (r *InMemoryEventArtistRepository) ArtistIDsForEvent(eventID string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var artistIDs []string

	for _, eventArtistAssociation := range r.eventArtistAssociations {
		if eventArtistAssociation.EventID == eventID {
			artistIDs = append(artistIDs, eventArtistAssociation.ArtistID)
		}
	}

	return artistIDs, nil
}

func (r *InMemoryEventArtistRepository) EventIDsForArtist(artistID string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var eventIDs []string

	for _, eventArtistAssociation := range r.eventArtistAssociations {
		if eventArtistAssociation.ArtistID == artistID {
			eventIDs = append(eventIDs, eventArtistAssociation.EventID)
		}
	}

	return eventIDs, nil
}

func (r *InMemoryEventArtistRepository) CreateEventArtistAssociation(eventArtistAssociation entities.EventArtist) (entities.EventArtist, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	eventArtistAssociation.ID = strconv.Itoa(len(r.eventArtistAssociations) + 1)
	r.eventArtistAssociations = append(r.eventArtistAssociations, eventArtistAssociation)
	return eventArtistAssociation, nil
}

func (r *InMemoryEventArtistRepository) DeleteEventArtistAssociation(id string) (entities.EventArtist, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, eventArtistAssociation := range r.eventArtistAssociations {
		if eventArtistAssociation.ID == id {
			r.eventArtistAssociations = append(r.eventArtistAssociations[:i], r.eventArtistAssociations[i+1:]...)
			return eventArtistAssociation, nil
		}
	}
	return entities.EventArtist{}, errors.New("event artist associations not found")
}

func (r *InMemoryEventArtistRepository) DeleteArtistFromItsEvents(artistID string) ([]entities.EventArtist, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var remainingAssociations []entities.EventArtist
	var deletedAssociations []entities.EventArtist
	for _, association := range r.eventArtistAssociations {
		if association.ArtistID == artistID {
			deletedAssociations = append(deletedAssociations, association)
		} else {
			remainingAssociations = append(remainingAssociations, association)
		}
	}
	r.eventArtistAssociations = remainingAssociations
	return deletedAssociations, nil
}

func (r *InMemoryEventArtistRepository) DeleteEventFromItsArtists(eventID string) ([]entities.EventArtist, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var remainingAssociations []entities.EventArtist
	var deletedAssociations []entities.EventArtist
	for _, association := range r.eventArtistAssociations {
		if association.EventID == eventID {
			deletedAssociations = append(deletedAssociations, association)
		} else {
			remainingAssociations = append(remainingAssociations, association)
		}
	}
	r.eventArtistAssociations = remainingAssociations
	return deletedAssociations, nil
}

func (r *InMemoryEventArtistRepository) UpdateEventArtistAssociation(id string, updatedEventArtistAssociation entities.EventArtist) (entities.EventArtist, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, eventArtistAssociation := range r.eventArtistAssociations {
		if eventArtistAssociation.ID == id {
			r.eventArtistAssociations[i] = updatedEventArtistAssociation
			r.eventArtistAssociations[i].ID = id
			return r.eventArtistAssociations[i], nil
		}
	}

	return entities.EventArtist{}, errors.New("event artist association not found")
}
