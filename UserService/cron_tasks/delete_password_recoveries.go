package cron_tasks

import (
	"time"

	"github.com/panicmilos/druz.io/UserService/models"
	"github.com/panicmilos/druz.io/UserService/services"

	"gorm.io/gorm"
)

func DeletePasswordRecoveryRequests() {
	db := services.Provider.Get(services.AppDatabaseInstance).(*gorm.DB)
	db.Where("expires_at < ?", time.Now()).Delete(&[]models.PasswordRecovery{})
}
