package entities

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Name        string `gorm:"column:name;unique;not null" json:"name" validate:"required,min=3,max=30"`
	Description string `gorm:"column:description" json:"description" validate:"max=256"`
}
