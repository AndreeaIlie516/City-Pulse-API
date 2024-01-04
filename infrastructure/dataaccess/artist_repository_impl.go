package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"errors"

	"gorm.io/gorm"
)

type GormArtistRepository struct {
	Db *gorm.DB
}

func NewGormArtistRepository(db *gorm.DB) *GormArtistRepository {
	return &GormArtistRepository{Db: db}
}

func (r *GormArtistRepository) AllArtists() ([]entities.Artist, error) {
	var artists []entities.Artist
	result := r.Db.Find(&artists)
	return artists, result.Error
}

func (r *GormArtistRepository) AllArtistIDs() ([]uint, error) {
	var artistIDs []uint

	if err := r.Db.Model(&entities.Artist{}).Select("ID").Find(&artistIDs).Error; err != nil {
		return nil, err
	}

	return artistIDs, nil
}

func (r *GormArtistRepository) ArtistByID(id uint) (*entities.Artist, error) {
	var artist entities.Artist

	if err := r.Db.First(&artist, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("artist not found")
		}
		return nil, err
	}

	return &artist, nil
}

func (r *GormArtistRepository) CreateArtist(artist entities.Artist) (entities.Artist, error) {
	if err := r.Db.Create(&artist).Error; err != nil {
		return entities.Artist{}, err
	}
	return artist, nil
}

func (r *GormArtistRepository) DeleteArtist(id uint) (entities.Artist, error) {
	var artist entities.Artist

	if err := r.Db.First(&artist, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Artist{}, errors.New("artist not found")
		}
		return entities.Artist{}, err
	}

	if err := r.Db.Delete(&artist).Error; err != nil {
		return entities.Artist{}, err
	}

	return artist, nil
}

func (r *GormArtistRepository) UpdateArtist(id uint, updatedArtist entities.Artist) (entities.Artist, error) {
	var artist entities.Artist

	if err := r.Db.First(&artist, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Artist{}, errors.New("artist not found")
		}
		return entities.Artist{}, err
	}

	if err := r.Db.Model(&artist).Updates(updatedArtist).Error; err != nil {
		return entities.Artist{}, err
	}

	return artist, nil
}
