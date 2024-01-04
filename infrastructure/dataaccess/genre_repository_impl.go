package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"errors"

	"gorm.io/gorm"
)

type GormGenreRepository struct {
	Db *gorm.DB
}

func NewGormGenreRepository(db *gorm.DB) *GormGenreRepository {
	return &GormGenreRepository{Db: db}
}

func (r *GormGenreRepository) AllGenres() ([]entities.Genre, error) {
	var genres []entities.Genre
	result := r.Db.Find(&genres)
	return genres, result.Error
}

func (r *GormGenreRepository) AllGenreIDs() ([]uint, error) {
	var genreIDs []uint

	if err := r.Db.Model(&entities.Genre{}).Select("ID").Find(&genreIDs).Error; err != nil {
		return nil, err
	}

	return genreIDs, nil
}

func (r *GormGenreRepository) GenreByID(id uint) (*entities.Genre, error) {
	var genre entities.Genre

	if err := r.Db.First(&genre, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("genre not found")
		}
		return nil, err
	}

	return &genre, nil
}

func (r *GormGenreRepository) CreateGenre(genre entities.Genre) (entities.Genre, error) {
	if err := r.Db.Create(&genre).Error; err != nil {
		return entities.Genre{}, err
	}
	return genre, nil
}

func (r *GormGenreRepository) DeleteGenre(id uint) (entities.Genre, error) {
	var genre entities.Genre

	if err := r.Db.First(&genre, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Genre{}, errors.New("genre not found")
		}
		return entities.Genre{}, err
	}

	if err := r.Db.Delete(&genre).Error; err != nil {
		return entities.Genre{}, err
	}

	return genre, nil
}

func (r *GormGenreRepository) UpdateGenre(id uint, updatedGenre entities.Genre) (entities.Genre, error) {
	var genre entities.Genre

	if err := r.Db.First(&genre, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Genre{}, errors.New("genre not found")
		}
		return entities.Genre{}, err
	}

	if err := r.Db.Model(&genre).Updates(updatedGenre).Error; err != nil {
		return entities.Genre{}, err
	}

	return genre, nil
}
