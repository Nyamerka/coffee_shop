package models

import "time"

type User struct {
	ID           uint64    `gorm:"primary_key"`
	Name         string    `gorm:"type:varchar(100)"`
	Email        string    `gorm:"type:varchar(100);unique"`
	PasswordHash string    `gorm:"type:varchar(255)"`
	RegisterDate time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
