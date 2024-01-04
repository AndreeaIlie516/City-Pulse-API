package entities

import (
	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model
	Name        string `gorm:"column:name;unique;not null" json:"name" validate:"required,min=3,max=30"`
	IsBand      bool   `gorm:"column:name;not null" json:"is_band" validate:"required,boolean"`
	BandMembers string `gorm:"column:band_members" json:"band_members" validate:"required,max=100"`
	Description string `gorm:"column:description" json:"description" validate:"max=256"`
}
