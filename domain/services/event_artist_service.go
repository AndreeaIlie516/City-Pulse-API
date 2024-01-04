package services

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/repositories"
	"errors"
	"fmt"
)

type EventArtistService struct {
	Repo         repositories.EventArtistRepository
	EventRepo    repositories.EventRepository
	ArtistRepo   repositories.ArtistRepository
	LocationRepo repositories.LocationRepository
	CityRepo     repositories.CityRepository
}

type EventArtistDetail struct {
	Association entities.EventArtist
	Event       entities.Event
	Artist      entities.Artist
}

type EventWithArtists struct {
	Event   entities.Event
	Artists []entities.Artist
}

type ArtistWithEvents struct {
	Artist entities.Artist
	Events []entities.Event
}

func (service *EventArtistService) AllEventArtistAssociations() ([]entities.EventArtist, error) {
	eventArtistAssociations, err := service.Repo.AllEventArtistAssociations()
	if err != nil {
		return nil, err
	}
	return eventArtistAssociations, nil
}

func (service *EventArtistService) EventArtistAssociationByID(idStr string) (*EventArtistDetail, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return nil, errors.New("invalid ID format")
	}

	eventArtistAssociation, err := service.Repo.EventArtistAssociationByID(id)
	if err != nil {
		return nil, err
	}

	event, err := service.EventRepo.EventByID(eventArtistAssociation.EventID)
	if err != nil {
		return nil, err
	}

	artist, err := service.ArtistRepo.ArtistByID(eventArtistAssociation.ArtistID)
	if err != nil {
		return nil, err
	}

	eventArtistDetail := &EventArtistDetail{
		Association: *eventArtistAssociation,
		Event:       *event,
		Artist:      *artist,
	}
	return eventArtistDetail, nil
}

func (service *EventArtistService) EventArtistAssociation(eventIDStr string, artistIDStr string) (*EventArtistDetail, error) {
	var eventID uint
	if _, err := fmt.Sscanf(eventIDStr, "%d", &eventID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	event, err := service.EventRepo.EventByID(eventID)
	if err != nil {
		return nil, err
	}

	var artistID uint
	if _, err := fmt.Sscanf(artistIDStr, "%d", &artistID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	artist, err := service.ArtistRepo.ArtistByID(artistID)
	if err != nil {
		return nil, err
	}

	eventArtistAssociation, err := service.Repo.EventArtistAssociation(eventID, artistID)
	if err != nil {
		return nil, err
	}

	eventArtistDetail := &EventArtistDetail{
		Association: *eventArtistAssociation,
		Event:       *event,
		Artist:      *artist,
	}
	return eventArtistDetail, nil
}

func (service *EventArtistService) EventWithArtists(eventIDStr string) (*EventWithArtists, error) {
	var eventID uint
	if _, err := fmt.Sscanf(eventIDStr, "%d", &eventID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	event, err := service.EventRepo.EventByID(eventID)
	if err != nil {
		return &EventWithArtists{}, err
	}

	artistIDs, err := service.Repo.ArtistIDsForEvent(eventID)
	var artists []entities.Artist

	for _, artistID := range artistIDs {
		artist, err := service.ArtistRepo.ArtistByID(artistID)
		if err != nil {
			return &EventWithArtists{}, err
		}
		artists = append(artists, *artist)
	}

	eventWithArtists := &EventWithArtists{
		Event:   *event,
		Artists: artists,
	}

	return eventWithArtists, nil
}

func (service *EventArtistService) ArtistWithEvents(artistIDStr string) (*ArtistWithEvents, error) {
	var artistID uint
	if _, err := fmt.Sscanf(artistIDStr, "%d", &artistID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	artist, err := service.ArtistRepo.ArtistByID(artistID)
	if err != nil {
		return &ArtistWithEvents{}, err
	}

	eventIDs, err := service.Repo.EventIDsForArtist(artistID)
	var events []entities.Event

	for _, eventID := range eventIDs {
		event, err := service.EventRepo.EventByID(eventID)
		if err != nil {
			return &ArtistWithEvents{}, err
		}
		events = append(events, *event)
	}

	artistWithEvents := &ArtistWithEvents{
		Artist: *artist,
		Events: events,
	}

	return artistWithEvents, nil
}

func (service *EventArtistService) CreateEventArtistAssociation(eventArtistAssociation entities.EventArtist) (entities.EventArtist, error) {
	_, err := service.EventRepo.EventByID(eventArtistAssociation.EventID)
	if err != nil {
		return entities.EventArtist{}, err
	}

	_, err = service.ArtistRepo.ArtistByID(eventArtistAssociation.ArtistID)
	if err != nil {
		return entities.EventArtist{}, err
	}

	eventArtistAssociation, err = service.Repo.CreateEventArtistAssociation(eventArtistAssociation)
	if err != nil {
		return entities.EventArtist{}, err
	}
	return eventArtistAssociation, nil
}

func (service *EventArtistService) DeleteEventArtistAssociation(idStr string) (entities.EventArtist, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.EventArtist{}, errors.New("invalid ID format")
	}

	eventArtistAssociation, err := service.Repo.DeleteEventArtistAssociation(id)
	if err != nil {
		return entities.EventArtist{}, err
	}
	return eventArtistAssociation, nil
}

func (service *EventArtistService) DeleteArtistFromItsEvents(artistIDStr string) ([]entities.EventArtist, error) {
	var artistID uint
	if _, err := fmt.Sscanf(artistIDStr, "%d", &artistID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	_, err := service.ArtistRepo.ArtistByID(artistID)
	if err != nil {
		return []entities.EventArtist{}, err
	}

	artistEventAssociation, err := service.Repo.DeleteArtistFromItsEvents(artistID)
	if err != nil {
		return []entities.EventArtist{}, err
	}
	return artistEventAssociation, nil
}

func (service *EventArtistService) DeleteEventFromItsArtists(eventIDStr string) ([]entities.EventArtist, error) {
	var eventID uint
	if _, err := fmt.Sscanf(eventIDStr, "%d", &eventID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	_, err := service.EventRepo.EventByID(eventID)
	if err != nil {
		return []entities.EventArtist{}, err
	}

	eventArtistAssociation, err := service.Repo.DeleteEventFromItsArtists(eventID)
	if err != nil {
		return []entities.EventArtist{}, err
	}
	return eventArtistAssociation, nil
}

func (service *EventArtistService) UpdateEventArtistAssociation(idStr string, eventArtistAssociation entities.EventArtist) (entities.EventArtist, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.EventArtist{}, errors.New("invalid ID format")
	}

	eventArtistAssociation, err := service.Repo.UpdateEventArtistAssociation(id, eventArtistAssociation)
	if err != nil {
		return entities.EventArtist{}, err
	}
	return eventArtistAssociation, nil
}
