package repository

import (
	"UserService/models"

	"gorm.io/gorm"
)

type EducationsCollection struct {
	DB *gorm.DB
}

func (educationsCollection *EducationsCollection) DeleteByProfileId(profileId uint) *[]models.Education {
	educations := &[]models.Education{}

	educationsCollection.DB.Where("profile_id = ?", profileId).Delete(educations)

	return educations
}
