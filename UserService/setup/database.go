package setup

import (
	"github.com/panicmilos/druz.io/UserService/models"
	"github.com/panicmilos/druz.io/UserService/services"

	"gorm.io/gorm"
)

func SetupDatabase() {
	db := services.Provider.Get(services.AppDatabaseInstance).(*gorm.DB)

	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Profile{})
	db.AutoMigrate(&models.LivePlace{})
	db.AutoMigrate(&models.WorkPlace{})
	db.AutoMigrate(&models.Education{})
	db.AutoMigrate(&models.Interes{})
	db.AutoMigrate(&models.UserReport{})
	db.AutoMigrate(&models.PasswordRecovery{})
	db.AutoMigrate(&models.UserReactivation{})
	db.AutoMigrate(&models.UserBlock{})
}
