package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string    `gorm:"default:''" json:"first_name" validate:"omitempty,min=2,max=50"`
	LastName  string    `gorm:"default:''" json:"last_name" validate:"omitempty,min=2,max=50"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email" validate:"required,email"`
	Username  string    `gorm:"uniqueIndex;not null" json:"username" validate:"required,min=3,max=30"`
	Password  string    `gorm:"not null" json:"password,omitempty" validate:"required,min=6"`
	LastLogin time.Time `gorm:"default:NULL" json:"last_login,omitempty"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
}
