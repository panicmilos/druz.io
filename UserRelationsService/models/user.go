package models

import (
	"gorm.io/gorm"
)

// swagger:model Profile
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Disabled  bool
}
