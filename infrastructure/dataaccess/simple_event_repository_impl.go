package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"errors"
	"gorm.io/gorm"
)

type GormSimpleEventRepository struct {
	Db *gorm.DB
}

func NewGormSimpleEventRepository(db *gorm.DB) *GormSimpleEventRepository {
	return &GormSimpleEventRepository{Db: db}
}

func (r *GormSimpleEventRepository) AllEvents() ([]entities.SimpleEvent, error) {
	var events []entities.SimpleEvent
	result := r.Db.Find(&events)
	return events, result.Error
}

func (r *GormSimpleEventRepository) AllEventIDs() ([]uint, error) {
	var eventIDs []uint

	if err := r.Db.Model(&entities.SimpleEvent{}).Select("ID").Find(&eventIDs).Error; err != nil {
		return nil, err
	}

	return eventIDs, nil
}

func (r *GormSimpleEventRepository) EventByID(id uint) (*entities.SimpleEvent, error) {
	var event entities.SimpleEvent

	if err := r.Db.First(&event, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("event not found")
		}
		return nil, err
	}

	return &event, nil
}

func (r *GormSimpleEventRepository) CreateEvent(event entities.SimpleEvent) (entities.SimpleEvent, error) {
	if err := r.Db.Create(&event).Error; err != nil {
		return entities.SimpleEvent{}, err
	}
	return event, nil
}

func (r *GormSimpleEventRepository) DeleteEvent(id uint) (entities.SimpleEvent, error) {
	var event entities.SimpleEvent

	if err := r.Db.First(&event, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.SimpleEvent{}, errors.New("event not found")
		}
		return entities.SimpleEvent{}, err
	}

	if err := r.Db.Delete(&event).Error; err != nil {
		return entities.SimpleEvent{}, err
	}

	return event, nil
}

func (r *GormSimpleEventRepository) UpdateEvent(id uint, updatedEvent entities.SimpleEvent) (entities.SimpleEvent, error) {
	var event entities.SimpleEvent

	if err := r.Db.First(&event, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.SimpleEvent{}, errors.New("event not found")
		}
		return entities.SimpleEvent{}, err
	}

	if err := r.Db.Model(&event).Updates(updatedEvent).Error; err != nil {
		return entities.SimpleEvent{}, err
	}

	return event, nil
}

func (r *GormSimpleEventRepository) FavouriteEvents() ([]entities.SimpleEvent, error) {
	var favouriteEvents []entities.SimpleEvent

	if err := r.Db.Where("is_favourite = true").Find(&favouriteEvents).Error; err != nil {
		return nil, err
	}

	return favouriteEvents, nil
}

func (r *GormSimpleEventRepository) PrivateEvents() ([]entities.SimpleEvent, error) {
	var favouriteEvents []entities.SimpleEvent

	if err := r.Db.Where("is_private = true").Find(&favouriteEvents).Error; err != nil {
		return nil, err
	}

	return favouriteEvents, nil
}

func (r *GormSimpleEventRepository) AddEventToFavourites(id uint) (entities.SimpleEvent, error) {
	var event entities.SimpleEvent

	if err := r.Db.First(&event, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.SimpleEvent{}, errors.New("event not found")
		}
		return entities.SimpleEvent{}, err
	}

	if err := r.Db.Model(&event).Update("is_favourite", true).Error; err != nil {
		return entities.SimpleEvent{}, err
	}

	return event, nil
}

func (r *GormSimpleEventRepository) DeleteEventFromFavourites(id uint) (entities.SimpleEvent, error) {
	var event entities.SimpleEvent

	if err := r.Db.First(&event, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.SimpleEvent{}, errors.New("event not found")
		}
		return entities.SimpleEvent{}, err
	}

	if err := r.Db.Model(&event).Update("is_favourite", false).Error; err != nil {
		return entities.SimpleEvent{}, err
	}

	return event, nil
}
