package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"column:username;unique;not null" json:"username" validate:"required,min=3,max=25"`
	Password    string `gorm:"column:password;not null" json:"password" validate:"required,min=8,max=40"`
	FirstName   string `gorm:"column:first_name;not null" json:"first_name" validate:"required,min=3"`
	LastName    string `gorm:"column:last_name;not null" json:"last_name" validate:"required,min=3"`
	Email       string `gorm:"column:email;unique;not null" json:"email" validate:"required,email"`
	PhoneNumber string `gorm:"column:phone_number;unique;not null" json:"phone_number" validate:"required,e164"`
	Address     string `gorm:"column:address" json:"address" validate:"max=100"`
}
