package entities

import (
	"gorm.io/gorm"
)

type EventArtist struct {
	gorm.Model
	EventID   uint   `gorm:"column:event_id;not null" json:"event_id" validate:"required,number"`
	ArtistID  uint   `gorm:"column:artist_id;not null" json:"artist_id" validate:"required,number"`
	StartTime string `gorm:"column:start_time" json:"start_time" validate:"required,max=30"`
	EndTime   string `gorm:"column:end_time" json:"end_time" validate:"required,max=30"`
}
