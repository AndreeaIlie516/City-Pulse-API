package services

import (
	"City-Pulse-API/domain/entities"
	"City-Pulse-API/domain/repositories"
	"errors"
	"fmt"
)

type FavouriteEventService struct {
	Repo      repositories.FavouriteEventRepository
	EventRepo repositories.EventRepository
	UserRepo  repositories.UserRepository
}

type FavouriteEventDetail struct {
	Association entities.FavouriteEvent
	Event       entities.Event
	User        entities.User
}

type EventWithUsers struct {
	Event entities.Event
	Users []entities.User
}

type UserWithEvents struct {
	User   entities.User
	Events []entities.Event
}

func (service *FavouriteEventService) AllFavouriteEventAssociations() ([]entities.FavouriteEvent, error) {
	favouriteEventAssociations, err := service.Repo.AllFavouriteEventAssociations()
	if err != nil {
		return nil, err
	}
	return favouriteEventAssociations, nil
}

func (service *FavouriteEventService) FavouriteEventAssociationByID(idStr string) (*FavouriteEventDetail, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return nil, errors.New("invalid ID format")
	}

	favouriteEventAssociation, err := service.Repo.FavouriteEventAssociationByID(id)
	if err != nil {
		return nil, err
	}

	event, err := service.EventRepo.EventByID(favouriteEventAssociation.EventID)
	if err != nil {
		return nil, err
	}

	user, err := service.UserRepo.UserByID(favouriteEventAssociation.UserID)
	if err != nil {
		return nil, err
	}

	favouriteEventDetail := &FavouriteEventDetail{
		Association: *favouriteEventAssociation,
		Event:       *event,
		User:        *user,
	}
	return favouriteEventDetail, nil
}

func (service *FavouriteEventService) FavouriteEventAssociation(eventIDStr string, userIDStr string) (*FavouriteEventDetail, error) {
	var eventID uint
	if _, err := fmt.Sscanf(eventIDStr, "%d", &eventID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	event, err := service.EventRepo.EventByID(eventID)
	if err != nil {
		return nil, err
	}

	var userID uint
	if _, err := fmt.Sscanf(userIDStr, "%d", &userID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	user, err := service.UserRepo.UserByID(userID)
	if err != nil {
		return nil, err
	}

	favouriteEventAssociation, err := service.Repo.FavouriteEventAssociation(eventID, userID)
	if err != nil {
		return nil, err
	}

	favouriteEventDetail := &FavouriteEventDetail{
		Association: *favouriteEventAssociation,
		Event:       *event,
		User:        *user,
	}
	return favouriteEventDetail, nil
}

func (service *FavouriteEventService) EventWithUsers(eventIDStr string) (*EventWithUsers, error) {
	var eventID uint
	if _, err := fmt.Sscanf(eventIDStr, "%d", &eventID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	event, err := service.EventRepo.EventByID(eventID)
	if err != nil {
		return &EventWithUsers{}, err
	}

	userIDs, err := service.Repo.UserIDsForEvent(eventID)
	var users []entities.User

	for _, userID := range userIDs {
		user, err := service.UserRepo.UserByID(userID)
		if err != nil {
			return &EventWithUsers{}, err
		}
		users = append(users, *user)
	}

	eventWithUsers := &EventWithUsers{
		Event: *event,
		Users: users,
	}

	return eventWithUsers, nil
}

func (service *FavouriteEventService) UserWithEvents(userIDStr string) (*UserWithEvents, error) {
	var userID uint
	if _, err := fmt.Sscanf(userIDStr, "%d", &userID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	user, err := service.UserRepo.UserByID(userID)
	if err != nil {
		return &UserWithEvents{}, err
	}

	eventIDs, err := service.Repo.EventIDsForUser(userID)
	var events []entities.Event

	for _, eventID := range eventIDs {
		event, err := service.EventRepo.EventByID(eventID)
		if err != nil {
			return &UserWithEvents{}, err
		}
		events = append(events, *event)
	}

	userWithEvents := &UserWithEvents{
		User:   *user,
		Events: events,
	}

	return userWithEvents, nil
}

func (service *FavouriteEventService) AddEventToFavourites(favouriteEventAssociation entities.FavouriteEvent) (entities.FavouriteEvent, error) {
	_, err := service.EventRepo.EventByID(favouriteEventAssociation.EventID)
	if err != nil {
		return entities.FavouriteEvent{}, err
	}

	_, err = service.UserRepo.UserByID(favouriteEventAssociation.UserID)
	if err != nil {
		return entities.FavouriteEvent{}, err
	}

	favouriteEventAssociation, err = service.Repo.AddEventToFavourites(favouriteEventAssociation)
	if err != nil {
		return entities.FavouriteEvent{}, err
	}
	return favouriteEventAssociation, nil
}

func (service *FavouriteEventService) DeleteEventFromFavourites(idStr string) (entities.FavouriteEvent, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.FavouriteEvent{}, errors.New("invalid ID format")
	}

	favouriteEventAssociation, err := service.Repo.DeleteEventFromFavourites(id)
	if err != nil {
		return entities.FavouriteEvent{}, err
	}
	return favouriteEventAssociation, nil
}

func (service *FavouriteEventService) DeleteUserFromItsEvents(userIDStr string) ([]entities.FavouriteEvent, error) {
	var userID uint
	if _, err := fmt.Sscanf(userIDStr, "%d", &userID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	_, err := service.UserRepo.UserByID(userID)
	if err != nil {
		return []entities.FavouriteEvent{}, err
	}

	favouriteEventAssociation, err := service.Repo.DeleteUserFromItsEvents(userID)
	if err != nil {
		return []entities.FavouriteEvent{}, err
	}
	return favouriteEventAssociation, nil
}

func (service *FavouriteEventService) DeleteEventFromItsUsers(eventIDStr string) ([]entities.FavouriteEvent, error) {
	var eventID uint
	if _, err := fmt.Sscanf(eventIDStr, "%d", &eventID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	_, err := service.EventRepo.EventByID(eventID)
	if err != nil {
		return []entities.FavouriteEvent{}, err
	}

	favouriteEventAssociation, err := service.Repo.DeleteEventFromItsUsers(eventID)
	if err != nil {
		return []entities.FavouriteEvent{}, err
	}
	return favouriteEventAssociation, nil
}
