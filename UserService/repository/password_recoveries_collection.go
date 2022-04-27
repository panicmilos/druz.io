package repository

import (
	"UserService/models"

	"gorm.io/gorm"
)

type PasswordRecoveriesCollection struct {
	DB *gorm.DB
}

func (passwordRecoveriesCollection *PasswordRecoveriesCollection) ReadByProfileId(profileId uint) *models.PasswordRecovery {
	passwordRecovery := &models.PasswordRecovery{}

	result := passwordRecoveriesCollection.DB.Where("profile_id = ?", profileId).First(passwordRecovery)
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
