package models

import (
	"gorm.io/gorm"
)

// swagger:model User
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Disabled  bool
	Image     string
}
