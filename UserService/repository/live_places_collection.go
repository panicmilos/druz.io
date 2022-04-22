package repository

import (
	"UserService/models"

	"gorm.io/gorm"
)

type LivePlacesCollection struct {
	DB *gorm.DB
}

func (livePlacesCollection *LivePlacesCollection) DeleteByProfileId(profileId uint) *[]models.LivePlace {
	livePlaces := &[]models.LivePlace{}

	livePlacesCollection.DB.Where("profile_id = ?", profileId).Delete(livePlaces)

	return livePlaces
}
