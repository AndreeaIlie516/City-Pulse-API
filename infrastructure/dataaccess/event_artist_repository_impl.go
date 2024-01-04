package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"errors"
	"gorm.io/gorm"
)

type GormEventArtistRepository struct {
	Db *gorm.DB
}

func NewGormEventArtistRepository(db *gorm.DB) *GormEventArtistRepository {
	return &GormEventArtistRepository{Db: db}
}

func (r *GormEventArtistRepository) AllEventArtistAssociations() ([]entities.EventArtist, error) {
	var eventArtistAssociations []entities.EventArtist
	result := r.Db.Find(&eventArtistAssociations)
	return eventArtistAssociations, result.Error
}

func (r *GormEventArtistRepository) AllEventArtistAssociationIDs() ([]uint, error) {
	var eventArtistAssociationIDs []uint

	if err := r.Db.Model(&entities.EventArtist{}).Select("ID").Find(&eventArtistAssociationIDs).Error; err != nil {
		return nil, err
	}

	return eventArtistAssociationIDs, nil
}

func (r *GormEventArtistRepository) EventArtistAssociationByID(id uint) (*entities.EventArtist, error) {
	var eventArtistAssociation entities.EventArtist

	if err := r.Db.First(&eventArtistAssociation, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("eventArtistAssociation not found")
		}
		return nil, err
	}

	return &eventArtistAssociation, nil
}

func (r *GormEventArtistRepository) EventArtistAssociation(eventID uint, artistID uint) (*entities.EventArtist, error) {
	var eventArtistAssociation entities.EventArtist

	if err := r.Db.First(&eventArtistAssociation, "event_id = ? AND artist_id = ?", eventID, artistID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("eventArtistAssociation not found")
		}
		return nil, err
	}

	return &eventArtistAssociation, nil
}

func (r *GormEventArtistRepository) ArtistIDsForEvent(eventID uint) ([]uint, error) {
	var artistIDs []uint

	if err := r.Db.Where("event_id = ?", eventID).Model(&entities.EventArtist{}).Select("ID").Find(&artistIDs).Error; err != nil {
		return nil, err
	}

	return artistIDs, nil
}

func (r *GormEventArtistRepository) EventIDsForArtist(artistID uint) ([]uint, error) {
	var eventIDs []uint

	if err := r.Db.Where("artist_id = ?", artistID).Model(&entities.EventArtist{}).Select("event_id").Find(&eventIDs).Error; err != nil {
		return nil, err
	}

	return eventIDs, nil
}

func (r *GormEventArtistRepository) CreateEventArtistAssociation(eventArtistAssociation entities.EventArtist) (entities.EventArtist, error) {
	if err := r.Db.Create(&eventArtistAssociation).Error; err != nil {
		return entities.EventArtist{}, err
	}
	return eventArtistAssociation, nil
}

func (r *GormEventArtistRepository) DeleteEventArtistAssociation(id uint) (entities.EventArtist, error) {
	var eventArtistAssociation entities.EventArtist

	if err := r.Db.First(&eventArtistAssociation, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.EventArtist{}, errors.New("eventArtistAssociation not found")
		}
		return entities.EventArtist{}, err
	}

	if err := r.Db.Delete(&eventArtistAssociation).Error; err != nil {
		return entities.EventArtist{}, err
	}

	return eventArtistAssociation, nil
}

func (r *GormEventArtistRepository) DeleteArtistFromItsEvents(artistID uint) ([]entities.EventArtist, error) {
	var eventArtistAssociations []entities.EventArtist
	var deletedAssociations []entities.EventArtist

	if err := r.Db.Where("artist_id = ?", artistID).Find(&eventArtistAssociations).Error; err != nil {
		return []entities.EventArtist{}, err
	}

	deletedAssociations = append(deletedAssociations, eventArtistAssociations...)

	if err := r.Db.Delete(&eventArtistAssociations, "artist_id = ?", artistID).Error; err != nil {
		return []entities.EventArtist{}, err
	}

	return deletedAssociations, nil
}

func (r *GormEventArtistRepository) DeleteEventFromItsArtists(eventID uint) ([]entities.EventArtist, error) {
	var eventArtistAssociations []entities.EventArtist
	var deletedAssociations []entities.EventArtist

	if err := r.Db.Where("event_id = ?", eventID).Find(&eventArtistAssociations).Error; err != nil {
		return []entities.EventArtist{}, err
	}

	deletedAssociations = append(deletedAssociations, eventArtistAssociations...)

	if err := r.Db.Delete(&eventArtistAssociations, "event_id = ?", eventID).Error; err != nil {
		return []entities.EventArtist{}, err
	}

	return deletedAssociations, nil
}

func (r *GormEventArtistRepository) UpdateEventArtistAssociation(id uint, updatedEventArtistAssociation entities.EventArtist) (entities.EventArtist, error) {
	var eventArtistAssociation entities.EventArtist

	if err := r.Db.First(&eventArtistAssociation, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.EventArtist{}, errors.New("eventArtistAssociation not found")
		}
		return entities.EventArtist{}, err
	}

	if err := r.Db.Model(&eventArtistAssociation).Updates(updatedEventArtistAssociation).Error; err != nil {
		return entities.EventArtist{}, err
	}

	return eventArtistAssociation, nil
}
