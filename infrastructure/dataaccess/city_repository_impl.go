package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"errors"

	"gorm.io/gorm"
)

type GormCityRepository struct {
	Db *gorm.DB
}

func NewGormCityRepository(db *gorm.DB) *GormCityRepository {
	return &GormCityRepository{Db: db}
}

func (r *GormCityRepository) AllCities() ([]entities.City, error) {
	var cities []entities.City
	result := r.Db.Find(&cities)
	return cities, result.Error
}

func (r *GormCityRepository) AllCityIDs() ([]uint, error) {
	var cityIDs []uint

	if err := r.Db.Model(&entities.City{}).Select("ID").Find(&cityIDs).Error; err != nil {
		return nil, err
	}

	return cityIDs, nil
}

func (r *GormCityRepository) CityByID(id uint) (*entities.City, error) {
	var city entities.City

	if err := r.Db.First(&city, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("city not found")
		}
		return nil, err
	}

	return &city, nil
}

func (r *GormCityRepository) CreateCity(city entities.City) (entities.City, error) {
	if err := r.Db.Create(&city).Error; err != nil {
		return entities.City{}, err
	}
	return city, nil
}

func (r *GormCityRepository) DeleteCity(id uint) (entities.City, error) {
	var city entities.City

	if err := r.Db.First(&city, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.City{}, errors.New("city not found")
		}
		return entities.City{}, err
	}

	if err := r.Db.Delete(&city).Error; err != nil {
		return entities.City{}, err
	}

	return city, nil
}

func (r *GormCityRepository) UpdateCity(id uint, updatedCity entities.City) (entities.City, error) {
	var city entities.City

	if err := r.Db.First(&city, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.City{}, errors.New("city not found")
		}
		return entities.City{}, err
	}

	if err := r.Db.Model(&city).Updates(updatedCity).Error; err != nil {
		return entities.City{}, err
	}

	return city, nil
}
