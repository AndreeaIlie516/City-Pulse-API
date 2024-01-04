package entities

import (
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	Name        string `gorm:"column:name;not null" json:"name" validate:"required,min=3,max=30"`
	Country     string `gorm:"column:country;not null" json:"country" validate:"required,min=3,max=30"`
	Description string `gorm:"column:description" json:"description" validate:"max=256"`
}
