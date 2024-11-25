package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string `gorm:"size:255;not null"`
	Email            string `gorm:"size:255;unique;not null"`
	PasswordHash     string `gorm:"not null"`
	RegistrationDate string `gorm:"not null"`
}
