package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// User our Struct
type User struct {
	gorm.Model
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password    string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}