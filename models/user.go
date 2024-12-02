package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string    `gorm:"not null" json:"first_name" validate:"required,min=2,max=50"`           // User's first name
	LastName  string    `gorm:"not null" json:"last_name" validate:"required,min=2,max=50"`            // User's last name
	Email     string    `gorm:"uniqueIndex;not null" json:"email" validate:"required,email"`           // Email must be unique and valid
	Username  string    `gorm:"uniqueIndex;not null" json:"username" validate:"required,min=3,max=30"` // Username must be unique
	Password  string    `gorm:"not null" json:"password,omitempty" validate:"required,min=6"`          // Password must have a minimum length
	LastLogin time.Time `gorm:"default:NULL" json:"last_login,omitempty"`                              // Tracks the last login time
	IsActive  bool      `gorm:"default:true" json:"is_active"`                                         // Indicates if the user is active
}
