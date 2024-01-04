package dataaccess

import (
	"City-Pulse-API/domain/entities"
	"errors"

	"gorm.io/gorm"
)

type GormArtistGenreRepository struct {
	Db *gorm.DB
}

func NewGormArtistGenreRepository(db *gorm.DB) *GormArtistGenreRepository {
	return &GormArtistGenreRepository{Db: db}
}

func (r *GormArtistGenreRepository) AllArtistGenreAssociations() ([]entities.ArtistGenre, error) {
	var artistGenreAssociations []entities.ArtistGenre
	result := r.Db.Find(&artistGenreAssociations)
	return artistGenreAssociations, result.Error
}

func (r *GormArtistGenreRepository) AllArtistGenreAssociationIDs() ([]uint, error) {
	var artistGenreAssociationIDs []uint

	if err := r.Db.Model(&entities.ArtistGenre{}).Select("ID").Find(&artistGenreAssociationIDs).Error; err != nil {
		return nil, err
	}

	return artistGenreAssociationIDs, nil
}

func (r *GormArtistGenreRepository) ArtistGenreAssociationByID(id uint) (*entities.ArtistGenre, error) {
	var artistGenreAssociation entities.ArtistGenre

	if err := r.Db.First(&artistGenreAssociation, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("artistGenreAssociation not found")
		}
		return nil, err
	}

	return &artistGenreAssociation, nil
}

func (r *GormArtistGenreRepository) ArtistGenreAssociation(artistID uint, genreID uint) (*entities.ArtistGenre, error) {
	var artistGenreAssociation entities.ArtistGenre

	if err := r.Db.First(&artistGenreAssociation, "artist_id = ? AND genre_id = ?", artistID, genreID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("artistGenreAssociation not found")
		}
		return nil, err
	}

	return &artistGenreAssociation, nil
}

func (r *GormArtistGenreRepository) GenreIDsForArtist(artistID uint) ([]uint, error) {
	var genreIDs []uint

	if err := r.Db.Where("artist_id = ?", artistID).Model(&entities.ArtistGenre{}).Select("genre_id").Find(&genreIDs).Error; err != nil {
		return nil, err
	}

	return genreIDs, nil
}

func (r *GormArtistGenreRepository) ArtistIDsForGenre(genreID uint) ([]uint, error) {
	var artistIDs []uint

	if err := r.Db.Where("genre_id = ?", genreID).Model(&entities.ArtistGenre{}).Select("artist_id").Find(&artistIDs).Error; err != nil {
		return nil, err
	}

	return artistIDs, nil
}

func (r *GormArtistGenreRepository) CreateArtistGenreAssociation(artistGenreAssociation entities.ArtistGenre) (entities.ArtistGenre, error) {
	if err := r.Db.Create(&artistGenreAssociation).Error; err != nil {
		return entities.ArtistGenre{}, err
	}
	return artistGenreAssociation, nil
}

func (r *GormArtistGenreRepository) DeleteArtistGenreAssociation(id uint) (entities.ArtistGenre, error) {
	var artistGenreAssociation entities.ArtistGenre

	if err := r.Db.First(&artistGenreAssociation, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.ArtistGenre{}, errors.New("artistGenreAssociation not found")
		}
		return entities.ArtistGenre{}, err
	}

	if err := r.Db.Delete(&artistGenreAssociation).Error; err != nil {
		return entities.ArtistGenre{}, err
	}

	return artistGenreAssociation, nil
}

func (r *GormArtistGenreRepository) DeleteGenreFromItsArtists(genreID uint) ([]entities.ArtistGenre, error) {
	var artistGenreAssociations []entities.ArtistGenre
	var deletedAssociations []entities.ArtistGenre

	if err := r.Db.Where("genreID = ?", genreID).Find(&artistGenreAssociations).Error; err != nil {
		return []entities.ArtistGenre{}, err
	}

	deletedAssociations = append(deletedAssociations, artistGenreAssociations...)

	if err := r.Db.Delete(&artistGenreAssociations, "genreID = ?", genreID).Error; err != nil {
		return []entities.ArtistGenre{}, err
	}

	return deletedAssociations, nil
}

func (r *GormArtistGenreRepository) DeleteArtistFromItsGenres(artistID uint) ([]entities.ArtistGenre, error) {
	var artistGenreAssociations []entities.ArtistGenre
	var deletedAssociations []entities.ArtistGenre

	if err := r.Db.Where("artistID = ?", artistID).Find(&artistGenreAssociations).Error; err != nil {
		return []entities.ArtistGenre{}, err
	}

	deletedAssociations = append(deletedAssociations, artistGenreAssociations...)

	if err := r.Db.Delete(&artistGenreAssociations, "artistID = ?", artistID).Error; err != nil {
		return []entities.ArtistGenre{}, err
	}

	return deletedAssociations, nil
}

func (r *GormArtistGenreRepository) UpdateArtistGenreAssociation(id uint, updatedArtistGenreAssociation entities.ArtistGenre) (entities.ArtistGenre, error) {
	var artistGenreAssociation entities.ArtistGenre

	if err := r.Db.First(&artistGenreAssociation, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.ArtistGenre{}, errors.New("artistGenreAssociation not found")
		}
		return entities.ArtistGenre{}, err
	}

	if err := r.Db.Model(&artistGenreAssociation).Updates(updatedArtistGenreAssociation).Error; err != nil {
		return entities.ArtistGenre{}, err
	}

	return artistGenreAssociation, nil
}
