package models

import (
	"time"

	
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string    `gorm:"uniqueIndex;not null" json:"email" validate:"required,email"`
	Username  string    `gorm:"uniqueIndex;not null" json:"username" validate:"required,min=3,max=30"`
	Password  string    `gorm:"not null" json:"-" validate:"required,min=6"`
	LastLogin time.Time `json:"last_login,omitempty"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
}
