package models

import (
	"time"

	"gorm.io/gorm"
)

// swagger:model Account
type Account struct {
	gorm.Model
	Email    string
	Password string
	Salt     string

	Profile Profile
}

// swagger:model Profile
type Profile struct {
	gorm.Model
	FirstName string
	LastName  string
	Birthday  time.Time
	Gender    Gender

	AccountID uint
}
