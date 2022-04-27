package repository

import (
	"UserService/models"

	"gorm.io/gorm"
)

type UserReportsCollection struct {
	DB *gorm.DB
}

func (userReportsCollection *UserReportsCollection) Create(report *models.UserReport) *models.UserReport {
	userReportsCollection.DB.Create(report)

	return report
}

func (userReportsCollection *UserReportsCollection) ReadById(id uint) *models.UserReport {
	report := &models.UserReport{}

	result := userReportsCollection.DB.First(report, id)
	if result.RowsAffected == 0 {
		return nil
	}

	return report
}

func (userReportsCollection *UserReportsCollection) Delete(id uint) *models.UserReport {
	report := &models.UserReport{}

	userReportsCollection.DB.Delete(report, id)

	return report
}
