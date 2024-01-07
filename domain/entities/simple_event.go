package entities

import "gorm.io/gorm"

type SimpleEvent struct {
	gorm.Model
	Time        string `gorm:"column:time" json:"time" validate:"required,max=30"`
	Band        string `gorm:"column:band;unique;not null" json:"band" validate:"required,bandValidator"`
	Location    string `gorm:"column:location;not null" json:"location" validate:"required,min=3,max=30"`
	ImageUrl    string `gorm:"column:image_url" json:"image_url" validate:"max=100"`
	Description string `gorm:"column:description" json:"description" validate:"max=256"`
	IsPrivate   bool   `gorm:"column:is_private;default:true" json:"is_private"`
	IsFavourite bool   `gorm:"column:is_favourite;default:false" json:"is_favourite"`
}
