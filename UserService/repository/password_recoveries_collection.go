package repository

import (
	"github.com/panicmilos/druz.io/UserService/models"

	"gorm.io/gorm"
)

type PasswordRecoveriesCollection struct {
	DB *gorm.DB
}

func (passwordRecoveriesCollection *PasswordRecoveriesCollection) ReadByProfileId(profileId uint) *models.PasswordRecovery {
	passwordRecovery := &models.PasswordRecovery{}

	query := passwordRecoveriesCollection.DB.Table("password_recoveries")
	query.Joins("JOIN profiles p ON password_recoveries.profile_id = p.id").Where("(p.disabled is NULL OR p.disabled = 0) AND p.deleted_at is NULL")
	result := query.Where("profile_id = ?", profileId).First(passwordRecovery)
	if result.RowsAffected == 0 {
		return nil
	}

	return passwordRecovery
}

func (passwordRecoveriesCollection *PasswordRecoveriesCollection) Create(passwordRecovery *models.PasswordRecovery) *models.PasswordRecovery {
	passwordRecoveriesCollection.DB.Create(passwordRecovery)

	return passwordRecovery
}

func (passwordRecoveriesCollection *PasswordRecoveriesCollection) DeleteByProfileId(profileId uint) *[]models.PasswordRecovery {
	passwordRecoveries := &[]models.PasswordRecovery{}

	passwordRecoveriesCollection.DB.Where("profile_id = ?", profileId).Delete(passwordRecoveries)

	return passwordRecoveries
}
