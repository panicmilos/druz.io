package models

import "gorm.io/gorm"

type UserReport struct {
	gorm.Model

	ReportedId uint
	Reported   Profile

	ReportedById uint
	ReportedBy   Profile

	Reason string
}
