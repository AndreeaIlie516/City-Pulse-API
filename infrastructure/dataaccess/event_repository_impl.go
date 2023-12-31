package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"errors"

	"gorm.io/gorm"
)

type GormEventRepository struct {
	Db *gorm.DB
}

func NewGormEventRepository(db *gorm.DB) *GormEventRepository {
	return &GormEventRepository{Db: db}
}

func (r *GormEventRepository) AllEvents() ([]entities.Event, error) {
	var events []entities.Event
	result := r.Db.Find(&events)
	return events, result.Error
}

func (r *GormEventRepository) AllEventIDs() ([]uint, error) {
	var eventIDs []uint

	if err := r.Db.Model(&entities.Event{}).Select("ID").Find(&eventIDs).Error; err != nil {
		return nil, err
	}

	return eventIDs, nil
}

func (r *GormEventRepository) EventByID(id uint) (*entities.Event, error) {
	var event entities.Event

	if err := r.Db.First(&event, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("event not found")
		}
		return nil, err
	}

	return &event, nil
}

func (r *GormEventRepository) EventIDsForLocation(locationID uint) ([]uint, error) {
	var eventIDs []uint

	if err := r.Db.Where("location_id = ?", locationID).Model(&entities.Event{}).Select("ID").Find(&eventIDs).Error; err != nil {
		return nil, err
	}

	return eventIDs, nil
}

func (r *GormEventRepository) CreateEvent(event entities.Event) (entities.Event, error) {
	if err := r.Db.Create(&event).Error; err != nil {
		return entities.Event{}, err
	}
	return event, nil
}

func (r *GormEventRepository) DeleteEvent(id uint) (entities.Event, error) {
	var event entities.Event

	if err := r.Db.First(&event, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Event{}, errors.New("event not found")
		}
		return entities.Event{}, err
	}

	if err := r.Db.Delete(&event).Error; err != nil {
		return entities.Event{}, err
	}

	return event, nil
}

func (r *GormEventRepository) UpdateEvent(id uint, updatedEvent entities.Event) (entities.Event, error) {
	var event entities.Event

	if err := r.Db.First(&event, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Event{}, errors.New("event not found")
		}
		return entities.Event{}, err
	}

	if err := r.Db.Model(&event).Updates(updatedEvent).Error; err != nil {
		return entities.Event{}, err
	}

	return event, nil
}

func (r *GormEventRepository) FavouriteEvents() ([]entities.Event, error) {
	var favouriteEvents []entities.Event

	if err := r.Db.Where("is_favourite = true").Find(&favouriteEvents).Error; err != nil {
		return nil, err
	}

	return favouriteEvents, nil
}

func (r *GormEventRepository) AddEventToFavourites(id uint) (entities.Event, error) {
	var event entities.Event

	if err := r.Db.First(&event, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Event{}, errors.New("event not found")
		}
		return entities.Event{}, err
	}

	if err := r.Db.Model(&event).Update("is_favourite", true).Error; err != nil {
		return entities.Event{}, err
	}

	return event, nil
}

func (r *GormEventRepository) DeleteEventFromFavourites(id uint) (entities.Event, error) {
	var event entities.Event

	if err := r.Db.First(&event, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Event{}, errors.New("event not found")
		}
		return entities.Event{}, err
	}

	if err := r.Db.Model(&event).Update("is_favourite", false).Error; err != nil {
		return entities.Event{}, err
	}

	return event, nil
}
