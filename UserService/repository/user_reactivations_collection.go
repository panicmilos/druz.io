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

	query := userReactivationsCollection.DB.Table("user_reactivations")
	query.Joins("JOIN profiles p ON user_reactivations.profile_id = p.id").Where("(p.disabled is NULL OR p.disabled = 0) AND p.deleted_at is NULL")
	result := query.Where("profile_id = ?", profileId).First(userReactivation)
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
