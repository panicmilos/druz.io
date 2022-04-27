package models

import "time"

type PasswordRecovery struct {
	ID        uint `gorm:"primarykey"`
	ProfileId uint
	Token     string
	ExpiresAt time.Time
}
