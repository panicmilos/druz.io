package models

import "gorm.io/gorm"

// swagger:model Account
type Account struct {
	gorm.Model
	Username string
	Password string
	Salt     string
}

// swagger:model Profile
type Profile struct {
	gorm.Model
	FirstName string
	LastName  string

	AccountId uint
	Account   Account
}
