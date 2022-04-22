package repository

import (
	"UserService/models"

	"gorm.io/gorm"
)

type WorkPlacesCollection struct {
	DB *gorm.DB
}

func (workPlacesCollection *WorkPlacesCollection) DeleteByProfileId(profileId uint) *[]models.WorkPlace {
	workPlaces := &[]models.WorkPlace{}

	workPlacesCollection.DB.Where("profile_id = ?", profileId).Delete(workPlaces)

	return workPlaces
}
