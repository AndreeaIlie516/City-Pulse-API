package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"column:username;unique;not null" json:"username" validate:"required,usernameValidator"`
	Password    string `gorm:"column:password;not null" json:"password" validate:"required,passwordValidator"`
	FirstName   string `gorm:"column:first_name;not null" json:"first_name" validate:"required,nameValidator"`
	LastName    string `gorm:"column:last_name;not null" json:"last_name" validate:"required,nameValidator"`
	Email       string `gorm:"column:email;unique;not null" json:"email" validate:"required,email"`
	PhoneNumber string `gorm:"column:phone_number;unique;not null" json:"phone_number" validate:"required,e164"`
	Address     string `gorm:"column:address" json:"address" validate:"max=100"`
}
