package models

import "gorm.io/gorm"

type UserReport struct {
	gorm.Model
	Reported   uint
	ReportedBy uint
	Reason     string
}
