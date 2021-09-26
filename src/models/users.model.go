package models

import (
	"gorm.io/gorm"
	"time"
)

// User our Struct
type User struct {
	gorm.Model
	Name      string `gorm:"size:255;not null;" json:"name"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	Password  string `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
