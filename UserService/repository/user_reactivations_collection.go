package repository

import (
	"github.com/panicmilos/druz.io/UserService/models"

	"gorm.io/gorm"
)

type UserReactivationsCollection struct {
	DB *gorm.DB
}

func (userReactivationsCollection *UserReactivationsCollection) ReadByProfileId(profileId uint) *models.UserReactivation {
	userReactivation := &models.UserReactivation{}

	result := userReactivationsCollection.DB.Where("profile_id = ?", profileId).First(userReactivation)
	if result.RowsAffected == 0 {
		return nil
	}

	return userReactivation
}

func (userReactivationsCollection *UserReactivationsCollection) Create(userReactivation *models.UserReactivation) *models.UserReactivation {
	userReactivationsCollection.DB.Create(userReactivation)

	return userReactivation
}

func (userReactivationsCollection *UserReactivationsCollection) DeleteByProfileId(profileId uint) *[]models.UserReactivation {
	userReactivations := &[]models.UserReactivation{}

	userReactivationsCollection.DB.Where("profile_id = ?", profileId).Delete(userReactivations)

	return userReactivations
}
