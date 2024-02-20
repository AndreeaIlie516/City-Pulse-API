package entities

import "gorm.io/gorm"

type FavouriteEvent struct {
	gorm.Model
	UserID  uint `gorm:"column:user_id;not null" json:"user_id" validate:"required,number"`
	EventID uint `gorm:"column:event_id;not null" json:"event_id" validate:"required,number"`
}
