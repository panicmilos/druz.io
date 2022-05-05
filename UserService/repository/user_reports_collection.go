package repository

import (
	"strings"

	"github.com/panicmilos/druz.io/UserService/dto"
	"github.com/panicmilos/druz.io/UserService/models"

	"gorm.io/gorm"
)

type UserReportsCollection struct {
	DB *gorm.DB
}

func (userReportsCollection *UserReportsCollection) Search(params *dto.UserReportsSearchParams) *[]models.UserReport {
	userReports := &[]models.UserReport{}

	query := userReportsCollection.DB.Table("user_reports")

	if len(strings.TrimSpace(params.Reason)) != 0 {
		query.Where("LOWER(user_reports.reason) like ?", "%"+strings.ToLower(params.Reason)+"%")
	}

	query.Joins("JOIN profiles p ON user_reports.reported = p.id").Where("(p.disabled is NULL OR p.disabled = 0) AND p.deleted_at is NULL")
	if len(strings.TrimSpace(params.Reported)) != 0 {
		query.Where("CONCAT(LOWER(p.first_name), ' ', LOWER(p.last_name)) like ?", "%"+strings.ToLower(params.Reported)+"%")
	}

	query.Joins("JOIN profiles p2 ON user_reports.reported_by = p2.id").Where("(p2.disabled is NULL OR p2.disabled = 0) AND p2.deleted_at is NULL")
	if len(strings.TrimSpace(params.ReportedBy)) != 0 {
		query.Where("CONCAT(LOWER(p2.first_name), ' ', LOWER(p2.last_name)) like ?", "%"+strings.ToLower(params.ReportedBy)+"%")
	}

	query.Find(userReports)
	return userReports
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
