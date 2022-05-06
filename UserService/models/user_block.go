package models

import "gorm.io/gorm"

type UserBlock struct {
	*gorm.Model

	BlockedById uint
	BlockedId   uint
}
