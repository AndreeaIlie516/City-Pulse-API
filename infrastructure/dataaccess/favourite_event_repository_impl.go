package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"errors"
	"gorm.io/gorm"
)

type GormFavouriteEventRepository struct {
	Db *gorm.DB
}

func (r *GormFavouriteEventRepository) AllFavouriteEventAssociationsIDs() ([]uint, error) {
	//TODO implement me
	panic("implement me")
}

func NewGormFavouriteEventRepository(db *gorm.DB) *GormFavouriteEventRepository {
	return &GormFavouriteEventRepository{Db: db}
}

func (r *GormFavouriteEventRepository) AllFavouriteEventAssociations() ([]entities.FavouriteEvent, error) {
	var favouriteEventAssociations []entities.FavouriteEvent
	result := r.Db.Find(&favouriteEventAssociations)
	return favouriteEventAssociations, result.Error
}

func (r *GormFavouriteEventRepository) AllFavouriteEventAssociationIDs() ([]uint, error) {
	var favouriteEventAssociationIDs []uint

	if err := r.Db.Model(&entities.FavouriteEvent{}).Select("ID").Find(&favouriteEventAssociationIDs).Error; err != nil {
		return nil, err
	}

	return favouriteEventAssociationIDs, nil
}

func (r *GormFavouriteEventRepository) FavouriteEventAssociationByID(id uint) (*entities.FavouriteEvent, error) {
	var favouriteEventAssociation entities.FavouriteEvent

	if err := r.Db.First(&favouriteEventAssociation, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("favouriteEventAssociation not found")
		}
		return nil, err
	}

	return &favouriteEventAssociation, nil
}

func (r *GormFavouriteEventRepository) FavouriteEventAssociation(eventID uint, userID uint) (*entities.FavouriteEvent, error) {
	var favouriteEventAssociation entities.FavouriteEvent

	if err := r.Db.First(&favouriteEventAssociation, "event_id = ? AND user_id = ?", eventID, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("favouriteEventAssociation not found")
		}
		return nil, err
	}

	return &favouriteEventAssociation, nil
}

func (r *GormFavouriteEventRepository) UserIDsForEvent(eventID uint) ([]uint, error) {
	var userIDs []uint

	if err := r.Db.Where("user_id = ?", eventID).Model(&entities.FavouriteEvent{}).Select("ID").Find(&userIDs).Error; err != nil {
		return nil, err
	}

	return userIDs, nil
}

func (r *GormFavouriteEventRepository) EventIDsForUser(userID uint) ([]uint, error) {
	var eventIDs []uint

	if err := r.Db.Where("user_id = ?", userID).Model(&entities.FavouriteEvent{}).Select("event_id").Find(&eventIDs).Error; err != nil {
		return nil, err
	}

	return eventIDs, nil
}

func (r *GormFavouriteEventRepository) AddEventToFavourites(favouriteEventAssociation entities.FavouriteEvent) (entities.FavouriteEvent, error) {
	if err := r.Db.Create(&favouriteEventAssociation).Error; err != nil {
		return entities.FavouriteEvent{}, err
	}
	return favouriteEventAssociation, nil
}

func (r *GormFavouriteEventRepository) DeleteEventFromFavourites(id uint) (entities.FavouriteEvent, error) {
	var favouriteEventAssociation entities.FavouriteEvent

	if err := r.Db.First(&favouriteEventAssociation, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.FavouriteEvent{}, errors.New("favouriteEventAssociation not found")
		}
		return entities.FavouriteEvent{}, err
	}

	if err := r.Db.Delete(&favouriteEventAssociation).Error; err != nil {
		return entities.FavouriteEvent{}, err
	}

	return favouriteEventAssociation, nil
}

func (r *GormFavouriteEventRepository) DeleteUserFromItsEvents(userID uint) ([]entities.FavouriteEvent, error) {
	var favouriteEventAssociations []entities.FavouriteEvent
	var deletedAssociations []entities.FavouriteEvent

	if err := r.Db.Where("user_id = ?", userID).Find(&favouriteEventAssociations).Error; err != nil {
		return []entities.FavouriteEvent{}, err
	}

	deletedAssociations = append(deletedAssociations, favouriteEventAssociations...)

	if err := r.Db.Delete(&favouriteEventAssociations, "user_id = ?", userID).Error; err != nil {
		return []entities.FavouriteEvent{}, err
	}

	return deletedAssociations, nil
}

func (r *GormFavouriteEventRepository) DeleteEventFromItsUsers(eventID uint) ([]entities.FavouriteEvent, error) {
	var favouriteEventAssociations []entities.FavouriteEvent
	var deletedAssociations []entities.FavouriteEvent

	if err := r.Db.Where("event_id = ?", eventID).Find(&favouriteEventAssociations).Error; err != nil {
		return []entities.FavouriteEvent{}, err
	}

	deletedAssociations = append(deletedAssociations, favouriteEventAssociations...)

	if err := r.Db.Delete(&favouriteEventAssociations, "event_id = ?", eventID).Error; err != nil {
		return []entities.FavouriteEvent{}, err
	}

	return deletedAssociations, nil
}
