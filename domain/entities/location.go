package entities

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	Name    string `gorm:"column:name;not null" json:"name" validate:"required,min=3,max=30"`
	CityID  uint   `gorm:"column:city_id;not null" json:"city_id" validate:"required,number"`
	Address string `gorm:"column:address" json:"address" validate:"max=100"`
}
