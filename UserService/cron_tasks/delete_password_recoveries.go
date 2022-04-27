package cron_tasks

import (
	"UserService/models"
	"UserService/services"
	"time"

	"gorm.io/gorm"
)

func DeletePasswordRecoveryRequests() {

	db := services.Provider.Get(services.AppDatabaseInstance).(*gorm.DB)
	db.Where("expires_at < ?", time.Now()).Delete(&[]models.PasswordRecovery{})
}
