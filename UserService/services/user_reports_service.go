package services

import (
	"UserService/dto"
	"UserService/errors"
	"UserService/models"
	"UserService/repository"
)

type UserReportsService struct {
	repository *repository.Repository
}

func (userReportsService *UserReportsService) Search(params *dto.UserReportsSearchParams) *[]models.UserReport {

	return userReportsService.repository.UserReportsCollection.Search(params)
}

func (userReportsService *UserReportsService) Create(report *models.UserReport) (*models.UserReport, error) {
	profile := userReportsService.repository.Users.ReadById(report.Reported)
	if profile == nil {
		return nil, errors.NewErrNotFound("Profile is not found")
	}

	if report.Reported == report.ReportedBy {
		return nil, errors.NewErrBadRequest("You can not report yourself")
	}

	return userReportsService.repository.UserReportsCollection.Create(report), nil
}

func (userReportsService *UserReportsService) Delete(id uint) (*models.UserReport, error) {
	report := userReportsService.repository.UserReportsCollection.ReadById(id)
	if report == nil {
		return nil, errors.NewErrNotFound("Report is not found")
	}

	return userReportsService.repository.UserReportsCollection.Delete(id), nil
}
