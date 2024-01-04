package entities

import (
	"gorm.io/gorm"
)

type ArtistGenre struct {
	gorm.Model
	ArtistID uint   `gorm:"column:artist_id;not null" json:"artist_id" validate:"required,number"`
	GenreID  uint   `gorm:"column:genre_id;not null" json:"genre_id" validate:"required,number"`
	Period   string `gorm:"column:period" json:"period" validate:"max=100"`
}
