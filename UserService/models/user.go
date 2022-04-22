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
	Role     Role

	Profile Profile
}

// swagger:model Profile
type Profile struct {
	gorm.Model
	FirstName string
	LastName  string
	Birthday  time.Time
	Gender    Gender

	About       string
	PhoneNumber string
	LivePlaces  []LivePlace
	WorkPlaces  []WorkPlace
	Educations  []Education
	Intereses   []Interes

	AccountID uint
}

type LivePlaceType int64

const (
	Currently LivePlaceType = iota
	Lived
	Birthplace
)

type LivePlace struct {
	ID            uint `gorm:"primarykey"`
	Place         string
	LivePlaceType LivePlaceType

	ProfileID uint
}

type WorkPlace struct {
	ID    uint `gorm:"primarykey"`
	Place string
	From  time.Time
	To    time.Time

	ProfileID uint
}

type Education struct {
	ID    uint `gorm:"primarykey"`
	Place string
	From  time.Time
	To    time.Time

	ProfileID uint
}

type Interes struct {
	ID      uint `gorm:"primarykey"`
	Interes string

	ProfileID uint
}
