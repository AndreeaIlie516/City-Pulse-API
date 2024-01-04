package entities

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name          string `gorm:"column:name;not null" json:"name" validate:"required,min=3,max=50"`
	LocationID    uint   `gorm:"column:location_id;not null" json:"location_id" validate:"required,number"`
	ImageUrl      string `gorm:"column:image_url" json:"image_url" validate:"max=100"`
	Date          string `gorm:"column:date;not null" json:"date" validate:"required,min=3,max=50"`
	StartTime     string `gorm:"column:start_time" json:"start_time" validate:"required,max=30"`
	EndTime       string `gorm:"column:end_time" json:"end_time" validate:"required,max=30"`
	OpenGatesTime string `gorm:"column:open_gates_time" json:"open_gates_time" validate:"required,max=30"`
	Description   string `gorm:"column:description" json:"description" validate:"max=256"`
	Type          string `gorm:"column:type" json:"type" validate:"max=100"`
	IsFavourite   bool   `gorm:"column:is_favourite;default:false" json:"is_favourite"`
}
