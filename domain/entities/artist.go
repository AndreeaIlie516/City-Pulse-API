package entities

import (
	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model
	Name        string `gorm:"column:name;unique;not null" json:"name" validate:"required,nameValidator"`
	IsBand      bool   `gorm:"column:is_band;not null" json:"is_band" validate:"required,boolean"`
	BandMembers string `gorm:"column:band_members" json:"band_members" validate:"required,bandValidator"`
	Description string `gorm:"column:description" json:"description" validate:"max=256"`
}
