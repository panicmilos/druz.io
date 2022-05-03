package models

import "time"

type UserReactivation struct {
	ID        uint `gorm:"primarykey"`
	ProfileId uint
	Token     string
	ExpiresAt time.Time
}
